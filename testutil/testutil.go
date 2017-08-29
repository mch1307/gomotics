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
	"github.com/mch1307/gomotics/nhc"
	"github.com/mch1307/gomotics/server"
)

const (
	connectHost  = "localhost"
	connectPort  = "8000"
	connectProto = "tcp"
)

var (
	actions = `{"cmd":"listactions","data":[{"id":0,"name":"light","type":1,"location":1,"value1":0},{"id":1,"name":"power switch","type":1,"location":2,"value1":0}]}
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
	failConf                                     = config.NhcConf{Host: "willFail", Port: 8000}
	testConf                                     = config.NhcConf{Host: "localhost", Port: 8000}
	command                                      = nhc.Event{ID: 1, Value: 100}
	myCmd                                        nhc.SimpleCmd
	fakeActionsMsg, fakeLocationsMsg, nhcMessage nhc.Message
	popFakeRun                                   bool
)

// PopFakeNhc populates in mem db with fake data for UT
func PopFakeNhc() {

	if !popFakeRun {
		json.Unmarshal([]byte(locations), &fakeLocationsMsg)
		nhc.Route(fakeLocationsMsg)
		json.Unmarshal([]byte(actions), &fakeActionsMsg)
		nhc.Route(fakeActionsMsg)
		nhc.BuildItems()
		popFakeRun = true
	}
}

func InitStubNHC() {
	config.Conf.NhcConfig.Host = connectHost
	config.Conf.NhcConfig.Port, _ = strconv.Atoi(connectPort)
	config.Conf.ServerConfig.ListenPort = 8081
	go MockNHC()
	go nhc.Listener()
	time.Sleep(500 * time.Millisecond)
	nhc.Init(&testConf)
	// call twice to test update items in persit.go
	nhc.Init(&testConf)
	/* 	myCmd.Cmd = "executeactions"
	   	myCmd.ID = 1
		   myCmd.Value = 100 */
	s := server.Server{}
	s.Initialize()
	go s.Run()
}

// MockNHC simulates a NHC controller on localhost:8000
func MockNHC() {
	l, err := net.Listen(connectProto, connectHost+":"+connectPort)
	if err != nil {
		os.Exit(1)
	}
	// Close the listener when the application closes.
	defer l.Close()
	fmt.Println("Listening on " + connectHost + ":" + connectPort)
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

type Sessions []*Session

var Clients Sessions

type Session struct {
	sType      string
	connection net.Conn
	reader     *bufio.Reader
	writer     *bufio.Writer
}

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

func (session *Session) Handle() {
	for {
		//fmt.Println("mock msg: ", nhcMessage.Cmd, nhcMessage.Event, nhcMessage.Data)
		message, _ := bufio.NewReader(session.connection).ReadBytes('\n')
		if len(message) > 0 {
			_ = json.Unmarshal(message, &nhcMessage)
			if nhcMessage.Cmd == "startevents" {
				fmt.Println("Listener session")
				session.sType = "listener"
				session.connection.Write([]byte(startEvents))
				//session.connection.Write([]byte(invalidMsg))
			} else if nhcMessage.Cmd == "listactions" {
				fmt.Println("Actions: ", nhcMessage.Cmd, nhcMessage.Event, session.sType)
				session.connection.Write([]byte(actions))
				nhcMessage.Cmd = "dropme"
			} else if nhcMessage.Cmd == "listlocations" {
				fmt.Println("Location: ", nhcMessage.Cmd, nhcMessage.Event, session.sType)
				session.connection.Write([]byte(locations))
				nhcMessage.Cmd = "dropme"
			} else if nhcMessage.Cmd == "executeactions" {
				fmt.Println("Event: ", nhcMessage.Cmd, nhcMessage.Event, session.sType)
				for _, cli := range Clients {
					if cli.sType == "listener" {
						cli.connection.Write([]byte(actionEvent))
					}
				}
			}
		}
	}
}
