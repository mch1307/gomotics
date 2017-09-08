// Package testutil pkg: set of utilities for facilitating tests
// Only used for unit/integration tests
//TODO: stub the UDP reply from NHC on port 10000
package testutil

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"strconv"
	"time"

	"github.com/mch1307/gomotics/config"
	"github.com/mch1307/gomotics/db"
	"github.com/mch1307/gomotics/nhc"
	"github.com/mch1307/gomotics/server"
	"github.com/mch1307/gomotics/types"
)

const (
	ConnectHost  = "localhost"
	ConnectPort  = "8000"
	connectProto = "tcp"
)

var (
	actions = `{"cmd":"listactions","data":[{"id":0,"name":"light","type":1,"location":1,"value1":0},{"id":1,"name":"power switch","type":1,"location":2,"value1":0},{"id":3,"name":"light","type":1,"location":1,"value1":0}]}
	`
	locations = `{"cmd":"listlocations","data":[{"id":0,"name":""},{"id":1,"name":"Living Room"},{"id":2,"name":"Kitchen"}]}
	`
	actionEvent = `{"event":"listactions","data":[{"id":1,"value1":100}]}
	`
	invalidMsg = `{"event":"listactions","data":[{"id":1,"value1":100}]
	`
	startEvents = `{"cmd":"startevents","data":[{"id":1,"value1":100}]}
	`
	invalidNHCMsg = `{"event":"listactions","data":[{"i":1,"val":100}]}
	`
	failConf = config.NhcConf{Host: "willFail", Port: 8000}
	testConf = config.NhcConf{Host: "localhost", Port: 8000}
	command  = types.Event{ID: 1, Value: 100}
	// MyCmd exported nhc.SimpleCmd
	MyCmd                                        = nhc.SimpleCmd{Cmd: "executeactions", ID: 1, Value: 100}
	fakeActionsMsg, fakeLocationsMsg, nhcMessage types.Message
	popFakeRun, initRun                          bool
	retries                                      = 0
)

// Sessions type used for managin session in NHC Stub (listener vs commands)
type Sessions []*Session

// Clients holds client session list
var Clients Sessions

// Session type holds the NHC stub client connection
type Session struct {
	sType      string
	connection net.Conn
	reader     *bufio.Reader
	writer     *bufio.Writer
}

// PopFakeNhc populates in mem db with fake data for UT
func PopFakeNhc() {

	if !popFakeRun {
		json.Unmarshal([]byte(locations), &fakeLocationsMsg)
		nhc.Route(&fakeLocationsMsg)
		json.Unmarshal([]byte(actions), &fakeActionsMsg)
		nhc.Route(&fakeActionsMsg)
		db.BuildItems()
		popFakeRun = true
	}
}

func IsTCPPortAvailable(port int) bool {
	if port < 1024 || port > 65500 {
		return false
	}
	conn, err := net.Listen("tcp", fmt.Sprintf("127.0.0.1:%d", port))
	if err != nil {
		return false
	}
	conn.Close()
	return true
}

// InitStubNHC initialize the NHC Stub and populates dummy data in mem 4 tests
func InitStubNHC() {

	//_, err := net.Dial("tcp", "127.0.0.1:8081")
	if IsTCPPortAvailable(8081) && IsTCPPortAvailable(8000) {
		fmt.Println("starting InitStubNHC")
		config.Conf.NhcConfig.Host = ConnectHost
		config.Conf.NhcConfig.Port, _ = strconv.Atoi(ConnectPort)
		config.Conf.ServerConfig.ListenPort = 8081
		config.Conf.ServerConfig.LogLevel = "DEBUG"
		go MockNHC()
		go nhc.Listener()
		time.Sleep(500 * time.Millisecond)
		//nhc.Init(&testConf)
		// call twice to test update items in persit.go
		//nhc.Init(&testConf)
		s := server.Server{}
		s.Initialize()
		go s.Run()
		//ws.Initialize()
		initRun = true
	} else {
		retries++
		if retries > 120 {
			return
		}
		fmt.Println("waiting for port to be available for the next test. Retries: ", retries)
		time.Sleep(time.Millisecond * 1000)
		InitStubNHC()
	}
}

// MockNHC simulates a NHC controller on localhost:8000
func MockNHC() {
	l, err := net.Listen(connectProto, ConnectHost+":"+ConnectPort)
	if err != nil {
		os.Exit(1)
	}
	// Close the listener when the application closes.
	defer l.Close()
	fmt.Println("Listening on " + ConnectHost + ":" + ConnectPort)
	for {
		// Listen for an incoming connection.
		conn, err := l.Accept()
		if err != nil {
			os.Exit(1)
		}
		// populate the list of Clients
		client := NewSession(conn)
		// handle connection in goroutine
		go client.Handle()
	}
}

// NewSession populates the Sessions (list of client connections) on client connect
func NewSession(conn net.Conn) *Session {
	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)
	session := &Session{
		connection: conn,
		reader:     reader,
		writer:     writer,
	}
	Clients = append(Clients, session)
	return session
}

// Handle handles the client connection
func (session *Session) Handle() {
	for {
		//fmt.Println("mock msg: ", nhcMessage.Cmd, nhcMessage.Event, nhcMessage.Data)
		message, _ := bufio.NewReader(session.connection).ReadBytes('\n')
		if len(message) > 0 {
			_ = json.Unmarshal(message, &nhcMessage)
			if nhcMessage.Cmd == "startevents" {
				//fmt.Println("Listener session")
				session.sType = "listener"
				session.connection.Write([]byte(startEvents))
				//session.connection.Write([]byte(invalidMsg))
			} else if nhcMessage.Cmd == "listactions" {
				//fmt.Println("Actions: ", nhcMessage.Cmd, nhcMessage.Event, session.sType)
				session.connection.Write([]byte(actions))
				nhcMessage.Cmd = "dropme"
			} else if nhcMessage.Cmd == "listlocations" {
				//fmt.Println("Location: ", nhcMessage.Cmd, nhcMessage.Event, session.sType)
				session.connection.Write([]byte(locations))
				nhcMessage.Cmd = "dropme"
			} else if nhcMessage.Cmd == "executeactions" {
				//fmt.Println("Event: ", nhcMessage.Cmd, nhcMessage.Event, session.sType)
				for _, cli := range Clients {
					if cli.sType == "listener" {
						cli.connection.Write([]byte(actionEvent))
					}
				}
			}
		}
	}
}
