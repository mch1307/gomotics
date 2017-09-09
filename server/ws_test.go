package server_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/gorilla/websocket"
	"github.com/mch1307/gomotics/db"
	"github.com/mch1307/gomotics/log"
	"github.com/mch1307/gomotics/testutil"
	"github.com/mch1307/gomotics/types"
)

var baseUrl string

func init() {
	fmt.Println("starting ws test")
	baseUrl = "http://" + testutil.ConnectHost + ":8081"
	testutil.InitStubNHC()
}

var origin = "http://localhost/"
var url = "ws://localhost:8081/events"

func wsDial(url string) (wsConn *websocket.Conn, ok bool, err error) {
	webS, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		fmt.Println("error connecting ws", err)
		return webS, false, err
	}
	//fmt.Println("websocket connect ok")
	return webS, true, nil
}

func Test_tWS(t *testing.T) {
	retry := 0
	ok := false
	ctl := 0
	var err error
	var wsConn *websocket.Conn
	tests := []struct {
		name       string
		id         int
		exName     string
		exLocation string
		exState    int
	}{
		{"action0", 3, "light", "Living Room", 0},
		{"action1", 1, "power switch", "Kitchen", 100},
	}
	//fmt.Println("# tests: ", len(tests))
	var msg types.Item
	time.Sleep(time.Second * 2)
	if wsConn, ok, err = wsDial(url); !ok {
		if retry < 11 {
			retry++
			fmt.Println("Retrying websocket connect due to error: ", err)
			fmt.Println("Attempt # ", retry)
			time.Sleep(time.Second * 1)
			Test_tWS(t)
		} else {
			fmt.Println("Could not connect after 10 attempts, err: ", err)
			return
		}
	}

	//time.Sleep(time.Second * 2)
	go func() {
		//defer ws.Close()
		//var tmp = make([]byte, 512)
		for {
			_, tmp, err := wsConn.ReadMessage()
			if err != nil {
				log.Error("read:", err)
				return
			}
			log.Error(string(tmp))
			err = json.Unmarshal(bytes.TrimSpace(tmp), &msg)
			if err != nil {
				log.Error("err parsing: ", err)
				log.Error(string(tmp))
			}
			//fmt.Println("ws reads ", msg)

		}
	}()

	time.Sleep(time.Second * 3)
	for _, tt := range tests {
		fmt.Println("start test ", tt.name)
		//ws.WriteMessage(websocket.PingMessage, nil)
		/* 		cmd := testutil.MyCmd
		   		cmd.ID = tt.id
		   		cmd.Value = tt.exState */
		//fmt.Println(cmd)
		time.Sleep(time.Millisecond * 500)
		var evt types.Event
		evt.ID = tt.id
		evt.Value = tt.exState
		db.ProcessEvent(evt)
		time.Sleep(time.Millisecond * 500)

		//fmt.Println("msg ", msg.ID)
		if msg.ID != tt.id || (msg.State != tt.exState) {
			t.Error("test failed  ", tt.name, tt.id, msg.ID, tt.exName, msg.Name, tt.exState, msg.State)
		}
		ctl++
	}
	defer wsConn.Close()
	//fmt.Println("tests ok: ", ctl)
}
