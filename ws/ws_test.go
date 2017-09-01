package ws_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/gorilla/websocket"
	"github.com/mch1307/gomotics/log"
	"github.com/mch1307/gomotics/nhc"
	"github.com/mch1307/gomotics/testutil"
	"github.com/mch1307/gomotics/types"
	//. "github.com/mch1307/gomotics/ws"
	//"golang.org/x/net/websocket"
)

var baseUrl string

//const healthMsg = `{"alive":true}`

func init() {
	baseUrl = "http://" + testutil.ConnectHost + ":8081"
	testutil.InitStubNHC()
}

var origin = "http://localhost/"
var url = "ws://localhost:8081/events"

func Test_tWS(t *testing.T) {
	retry := 0
	tests := []struct {
		name       string
		id         int
		exName     string
		exLocation string
		exState    int
	}{
		{"action0", 0, "light", "Living Room", 0},
		{"action1", 1, "power switch", "Kitchen", 100},
	}
	var msg types.Item
	ws, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		if retry == 0 {
			testutil.InitStubNHC()
		} else {
			fmt.Println(err)
		}
	}
	go func() {
		//defer ws.Close()
		//var tmp = make([]byte, 512)
		for {
			_, tmp, err := ws.ReadMessage()
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
			log.Debug("ws reads ", msg)

		}
	}()

	time.Sleep(time.Second * 1)
	for _, tt := range tests {
		fmt.Println("start test ", tt.name)
		//ws.WriteMessage(websocket.PingMessage, nil)
		cmd := testutil.MyCmd
		cmd.ID = tt.id
		cmd.Value = tt.exState
		fmt.Println(cmd)
		time.Sleep(time.Millisecond * 500)
		nhc.SendCommand(cmd.Stringify())
		fmt.Println("sending: ", cmd.ID)
		time.Sleep(time.Millisecond * 2000)

		fmt.Println("msg ", msg.ID)
		if msg.ID == tt.id {
			if msg.State != tt.exState {
				fmt.Println("testing...")
				t.Error("test failed  ", tt.name, tt.id, msg.ID, tt.exName, msg.Name, tt.exState, msg.State)
			}
		}
		/* 		if msg.ID == 1 {
			fmt.Println("abnormal connection")
			ws.WriteMessage(websocket.CloseAbnormalClosure, nil)
		} */
	}

}
