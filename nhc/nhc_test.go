package nhc_test

import (
	"reflect"
	"testing"
	"time"

	"github.com/mch1307/gomotics/db"
	"github.com/mch1307/gomotics/nhc"
	"github.com/mch1307/gomotics/testutil"
	"github.com/mch1307/gomotics/types"

	"github.com/mch1307/gomotics/config"
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
	failConf = config.NhcConf{Host: "willFail", Port: 8000}
	testConf = config.NhcConf{Host: "localhost", Port: 8000}
	command  = types.Event{ID: 1, Value: 100}
	//myCmd    nhc.SimpleCmd
)

func init() {
	// test failures
	//go Init(&failConf)

	testutil.InitStubNHC()
	//SendCommand(myCmd.Stringify())
}

func Test_getLocation(t *testing.T) {
	id := 1
	expect := "Living Room"
	t.Run("location", func(t *testing.T) {
		if got := db.GetLocation(id); !reflect.DeepEqual(got.Name, expect) {
			t.Errorf("getLocation() = %v, expected %v", got.Name, expect)
		}
	})

}

func TestGetAction(t *testing.T) {

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
			got := db.GetAction(tt.arg)
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
	nhc.SendCommand(testutil.MyCmd.Stringify())
	time.Sleep(300 * time.Millisecond)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := db.GetItems()
			for _, item := range got {
				if item.ID == tt.id {
					if item.State != tt.exVal {
						t.Errorf("GetItems() check item has proper status. Expected: %v, got: %v", tt.exVal, item.State)
					}
				}
			}
		})
	}
}
