package server

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net"
	"strconv"
	"time"

	"github.com/mch1307/gomotics/config"
	"github.com/mch1307/gomotics/db"
	"github.com/mch1307/gomotics/log"
	"github.com/mch1307/gomotics/types"
)

const (
	// SystemInfo holds NHC startevents
	SystemInfo = "{\"cmd\":\"systeminfo\"}"
	// RegisterCMD holds NHC startevents
	RegisterCMD = "{\"cmd\":\"startevents\"}"
	// ListActions holds NHC listactions
	ListActions = "{\"cmd\":\"listactions\"}"
	// ListLocations holds NHC listlocations
	ListLocations = "{\"cmd\":\"listlocations\"}"
	// ListEnergies holds NHC listenergy
	ListEnergies = "{\"cmd\":\"listenergy\"}"
	// ListThermostats holds NHC listthermostat
	ListThermostats = "{\"cmd\":\"listthermostat\"}"
)

var (
	nhcActions   []types.Action
	nhcLocations []types.Location
	nhcEvent     []types.Event
	nhcInfo      types.NHCSystemInfo
	nhcMessage   types.Message
)

// SimpleCmd type holding a nhc command
type SimpleCmd struct {
	Cmd   string `json:"cmd"`
	ID    int    `json:"id"`
	Value int    `json:"value1"`
}

// Stringify return the string version of SimpleCmd
func (sc SimpleCmd) Stringify() string {
	tmp, _ := json.Marshal(sc)
	return string(tmp)
}

// NhcInit sends list commands to NHC in order to get all equipments
func NhcInit(cfg *config.NhcConf) {
	if len(cfg.Host) == 0 {
		log.Debug("no nhc host in conf, starting discovery")
		nhcHost := Discover()
		log.Debug("discover returned ", nhcHost)
		config.Conf.NhcConfig.Host = nhcHost.String()
		cfg.Host = config.Conf.NhcConfig.Host
	}
	if cfg.Port == 0 {
		config.Conf.NhcConfig.Port = 8000
		cfg.Port = config.Conf.NhcConfig.Port

	}
	conn, err := ConnectNhc(cfg)
	if err != nil {
		log.Fatalf("Unable to connect to NHC host: %v. Error: %v", cfg.Host, err)
	}
	reader := json.NewDecoder(conn)
	fmt.Println("Connected to NHC unit: ", cfg.Host)
	log.Info("Connected to NHC unit: ", cfg.Host)

	// sends systeminfo command to NHC
	fmt.Fprintf(conn, SystemInfo+"\n")
	if err := reader.Decode(&nhcMessage); err != nil {
		log.Fatalf("Unable to parse NHC SystemInfo message: %v", err)
	}
	Route(&nhcMessage)

	// sends listlocations command to NHC
	fmt.Fprintf(conn, ListLocations+"\n")
	if err := reader.Decode(&nhcMessage); err != nil {
		log.Fatalf("Unable to parse NHC ListLocations message: %v", err)
	}
	Route(&nhcMessage)

	// sends listActions command to NHC
	fmt.Fprintf(conn, ListActions+"\n")
	if err := reader.Decode(&nhcMessage); err != nil {
		log.Fatalf("Unable to parse NHC ListActions message: %v", err)
	}
	Route(&nhcMessage)

	defer conn.Close()
	// Build the nhc collection
	db.BuildItems()
	if config.Conf.ServerConfig.LogLevel == "DEBUG" {
		db.Dump()
	}
	log.Info("Nhc init done")
}

// NhcListener start a connection to nhc host, register itself
// to receive and route all messages from nhc broadcast
func NhcListener() {
	var nhcMessage types.Message

	conn, err := ConnectNhc(&config.Conf.NhcConfig)
	if err != nil {
		log.Fatal("Fatal error connecting to NHC: ", err)
	}

	fmt.Fprintf(conn, RegisterCMD+"\n")

	for {
		reader := json.NewDecoder(conn)
		if err := reader.Decode(&nhcMessage); err != nil {
			log.Errorf("error decoding NHC message %v", err)
		}
		if nhcMessage.Cmd == "startevents" {
			log.Info("listener registered")
			nhcMessage.Cmd = "dropme"
		} else {
			log.Debug("received ", &nhcMessage.Cmd)
			Route(&nhcMessage)
		}
	}
}

// Route parse and route incoming message the right handler
func Route(msg *types.Message) {
	if msg.Cmd == "listlocations" {
		/* Commented err handling as all incoming msg have already been json parsed once */
		/* 		if err := json.Unmarshal(msg.Data, &nhcLocations); err != nil {
			log.Fatal(err)
		} */
		_ = json.Unmarshal(msg.Data, &nhcLocations)
		for idx := range nhcLocations {
			db.SaveLocation(nhcLocations[idx])
			//db.SaveItem(nhcLocations[idx])
		}
	} else if msg.Cmd == "listactions" {
		/* 		if err := json.Unmarshal(msg.Data, &nhcActions); err != nil {
			log.Fatal(err)
		} */
		_ = json.Unmarshal(msg.Data, &nhcActions)
		for idx := range nhcActions {
			db.SaveAction(nhcActions[idx])
			//db.SaveItem(nhcActions[idx])
		}
	} else if msg.Event == "listactions" {
		/* 		if err := json.Unmarshal(msg.Data, &nhcEvent); err != nil {
			log.Errorf("unable to parse message %v, err: %v", msg.Data, err)
		} */
		msg.Event = "dropme"
		_ = json.Unmarshal(msg.Data, &nhcEvent)
		for _, rec := range nhcEvent {
			WSPool.Broadcast <- db.ProcessEvent(rec)
			//db.SaveItem(nhcEvent[idx])
		}
	} else if msg.Cmd == "systeminfo" {
		msg.Cmd = "dropme"
		_ = json.Unmarshal(msg.Data, &nhcInfo)
		db.SaveNhcSysInfo(nhcInfo)
	}
}

// Discover discover NHC controller by sending UDP pkg on port 10000
// return NHC IP address and boolean
func Discover() net.IP {
	//	var err error
	var nhcConnectString net.IP
	var targetAddr *net.UDPAddr
	data, _ := hex.DecodeString("44")
	addr := net.UDPAddr{IP: net.IPv4bcast, Port: 10000}

	conn, err := net.ListenUDP("udp4", &net.UDPAddr{IP: GetOutboundIP(), Port: 18043})
	if err != nil {
		fmt.Println("err connect: ", err)
	}
	//defer conn.Close()
	_, err = conn.WriteToUDP(data, &addr)

	b := make([]byte, 1024)
	// goroutine for reading broadcast result
	go func() {
		for {
			defer conn.Close()
			_, targetAddr, err = conn.ReadFromUDP(b)
			if err != nil {
				log.Warnf("Error: UDP read error: %v", err)
				continue
			}
			// test "nhc" connection to replying IP to make sure targetAddr is a NHC controller
			connectString := net.TCPAddr{IP: targetAddr.IP, Port: 8000}
			if err == nil {
				_, err := net.DialTCP("tcp", nil, &connectString)
				//defer nhConn.Close()
				if err == nil {
					nhcConnectString = connectString.IP
					log.Debug("return IP: ", string(nhcConnectString))
					return
				}
			}
			defer conn.Close()
			//			return
		}
	}()
	time.Sleep(time.Second * 6)
	return nhcConnectString
}

//ConnectNhc establish connection to Niko Home Control IP module
func ConnectNhc(cfg *config.NhcConf) (conn *net.TCPConn, err error) {

	connectString, err := net.ResolveTCPAddr("tcp", cfg.Host+":"+strconv.Itoa(cfg.Port))
	if err != nil {
		log.Error("connNhc ", err)
		return nil, err
	}

	conn, err = net.DialTCP("tcp", nil, connectString)
	if err != nil {
		log.Fatal("error connecting to nhc: ", err)
	}
	return conn, err
}

// SendCommand send passed command to nhc
func SendCommand(cmd string) error {
	conn, _ := ConnectNhc(&config.Conf.NhcConfig)
	// no error handling as connect will exit in case of issue
	log.Debug("received command: ", cmd)
	fmt.Fprintf(conn, cmd+"\n")
	return nil
}
