package nhc

import (
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
	CONN_HOST = "ws2"
	CONN_PORT = "8000"
	CONN_TYPE = "tcp"
)

var (
	actions = `{"cmd":"listactions","data":[{"id":0,"name":"light","type":1,"location":1,"value1":0},{"id":1,"name":"power switch","type":1,"location":2,"value1":0}]}
	`
	locations = `{"cmd":"listlocations","data":[{"id":0,"name":""},{"id":1,"name":"Living Room"},{"id":2,"name":"Kitchen"}]}
	`
	actionEvent = `{"event":"listactions","data":[{"id":1,"value1":100}]}
	`
	fakeActionsMsg, fakeLocationsMsg Message
	popFakeRun                       bool
)

var testConf = config.NhcConf{Host: "ws2", Port: 8000}
var command = Event{ID: 1, Value: 100}

/* type nhCmdMessage struct {
	cmd string `json:"cmd"`
} */

func init() {
	config.Conf.NhcConfig.Host = "ws2"
	config.Conf.NhcConfig.Port = 8000
	go MockNHC()
	time.Sleep(1 * time.Second)
	//go Listener()
	time.Sleep(1 * time.Second)
	Init(&testConf)

}

func popFakeData() {

	if !popFakeRun {
		fmt.Println("popFake false")
		if err := json.Unmarshal([]byte(locations), &fakeLocationsMsg); err != nil {
			fmt.Println("Error unmarshalling location")
			panic(err)
		}
		Route(fakeLocationsMsg)

		if err := json.Unmarshal([]byte(actions), &fakeActionsMsg); err != nil {
			fmt.Println("Error unmarshalling action")
			panic(err)
		}
		Route(fakeActionsMsg)
		BuildItems()
		popFakeRun = true
	}
}

// MockNHC simulates a NHC controller on localhost:8000
func MockNHC() {
	l, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	// Close the listener when the application closes.
	defer l.Close()
	fmt.Println("Listening on " + CONN_HOST + ":" + CONN_PORT)
	for {
		// Listen for an incoming connection.
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		// Handle connections in a new goroutine.
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	var nhcMessage Message

	reader := json.NewDecoder(conn)
	for {
		//fmt.Println("mock msg: ", nhcMessage.Cmd, nhcMessage.Event, nhcMessage.Data)
		if err := reader.Decode(&nhcMessage); err != nil {
			fmt.Println("error reading input ", err)
		}
		if nhcMessage.Cmd == "listactions" {
			fmt.Println("Actions: ", nhcMessage.Cmd, nhcMessage.Event, nhcMessage.Data)
			conn.Write([]byte(actions))
			nhcMessage.Cmd = "dropme"
		} else if nhcMessage.Cmd == "listlocations" {
			fmt.Println("Location: ", nhcMessage.Cmd, nhcMessage.Event, nhcMessage.Data)
			conn.Write([]byte(locations))
			nhcMessage.Cmd = "dropme"
		} else if nhcMessage.Event == "listactions" {
			fmt.Println("Event: ", nhcMessage.Cmd, nhcMessage.Event, nhcMessage.Data)
			conn.Write([]byte(actionEvent))
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
	//popFakeData()
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

	ProcessEvent(command)
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
