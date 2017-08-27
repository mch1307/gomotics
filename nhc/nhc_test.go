package nhc

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/mch1307/gomotics/config"
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
	testConf = config.NhcConf{Host: "localhost", Port: 8000}
	command  = Event{ID: 1, Value: 100}
	myCmd    SimpleCmd
)

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
			if err := json.Unmarshal(message, &nhcMessage); err != nil {
				fmt.Println("error reading input ", err)
			}
			if nhcMessage.Cmd == "startevents" {
				fmt.Println("Listener session")
				session.sType = "listener"
				session.connection.Write([]byte(startEvents))
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

func init() {
	config.Conf.NhcConfig.Host = "localhost"
	config.Conf.NhcConfig.Port = 8000
	go MockNHC()

	go Listener()
	time.Sleep(500 * time.Millisecond)
	Init(&testConf)
	// call twice to test update items in persit.go
	Init(&testConf)
	myCmd.Cmd = "executeactions"
	myCmd.ID = 1
	myCmd.Value = 100
	//SendCommand(myCmd.Stringify())
}

// MockNHC simulates a NHC controller on localhost:8000
func MockNHC() {
	l, err := net.Listen(connectProto, connectHost+":"+connectPort)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	// Close the listener when the application closes.
	defer l.Close()
	fmt.Println("Listening on " + connectHost + ":" + connectPort)
	for {
		// Listen for an incoming connection.
		conn, err := l.Accept()

		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		// populate the list of Clients
		client := NewSession(conn)
		// handle connection in goroutine
		go client.Handle()
	}
}

func handleConnection(conn net.Conn) {
	var nhcMessage Message

	for {
		//fmt.Println("mock msg: ", nhcMessage.Cmd, nhcMessage.Event, nhcMessage.Data)
		message, _ := bufio.NewReader(conn).ReadBytes('\n')
		if len(message) > 0 {
			if err := json.Unmarshal(message, &nhcMessage); err != nil {
				fmt.Println("error reading input ", err)
			}
			if nhcMessage.Cmd == "listactions" {
				//fmt.Println("Actions: ", nhcMessage.Cmd, nhcMessage.Event, nhcMessage.Data)
				conn.Write([]byte(actions))
				nhcMessage.Cmd = "dropme"
			} else if nhcMessage.Cmd == "listlocations" {
				//fmt.Println("Location: ", nhcMessage.Cmd, nhcMessage.Event, nhcMessage.Data)
				conn.Write([]byte(locations))
				nhcMessage.Cmd = "dropme"
			} else if nhcMessage.Cmd == "executeactions" {
				//fmt.Println("Event: ", nhcMessage.Cmd, nhcMessage.Event, nhcMessage.Data)
				conn.Write([]byte(actionEvent))
			}
		}
	}
}
func Test_getLocation(t *testing.T) {
	//popFakeData()
	//Init(&testConf)
	id := 1
	expect := "Living Room"
	t.Run("location", func(t *testing.T) {
		if got := GetLocation(id); !reflect.DeepEqual(got.Name, expect) {
			t.Errorf("getLocation() = %v, expected %v", got.Name, expect)
		}
	})

}

func TestGetAction(t *testing.T) {

	type args struct {
		id int
	}
	tests := []struct {
		name       string
		arg        int
		exName     string
		exLocation string
	}{
		{"action0", 0, "light", "Living Room"},
		{"action1", 1, "power switch", "Kitchen"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetAction(tt.arg)
			if !reflect.DeepEqual(got.Name, tt.exName) {
				t.Errorf("GetAction() name = %v, want %v", got.Name, tt.exName)
			} else if !reflect.DeepEqual(got.Name, tt.exName) {
				t.Errorf("GetAction() location = %v, expect %v", got.Location, tt.exLocation)
			}
		})
	}
}

func TestGetItems(t *testing.T) {
	tests := []struct {
		name  string
		id    int
		exVal int
	}{
		{"fakeSwitch", 1, 100},
	}
	SendCommand(myCmd.Stringify())
	time.Sleep(100 * time.Millisecond)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetItems()
			for _, item := range got {
				if item.ID == tt.id {
					if item.State != 100 {
						t.Errorf("GetItems() check item has proper status. Expected: %v, got: %v", tt.exVal, item.State)
					}
				}
			}
		})
	}
}
