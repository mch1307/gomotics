package db_test

import (
	"reflect"
	"testing"
	"time"

	. "github.com/mch1307/gomotics/db"
	"github.com/mch1307/gomotics/nhc"
	"github.com/mch1307/gomotics/testutil"
)

func init() {
	// test failures
	//go Init(&failConf)
	testutil.InitStubNHC()
	//SendCommand(myCmd.Stringify())
}
func TestGetLocation(t *testing.T) {
	id := 1
	expect := "Living Room"
	t.Run("location", func(t *testing.T) {
		if got := GetLocation(id); !reflect.DeepEqual(got.Name, expect) {
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
	nhc.SendCommand(testutil.MyCmd.Stringify())
	time.Sleep(300 * time.Millisecond)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetItems()
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
