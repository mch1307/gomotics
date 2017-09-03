package db_test

import (
	"fmt"
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

	fmt.Println("starting db test")
	testutil.InitStubNHC()
	Dump()
	//SendCommand(myCmd.Stringify())
}
func TestGetLocation(t *testing.T) {
	id := 2
	//expect := "Living Room"
	expect := "Kitchen"
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
		{"fakeSwitch", 1, 0},
	}
	nhc.SendCommand(testutil.MyCmd.Stringify())
	time.Sleep(300 * time.Millisecond)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetItems()
			fmt.Println(len(got))
			Dump()
			for _, item := range got {
				if item.ID == tt.id {
					fmt.Println("ok got ", item.Name, " ", item.ID)
					if item.State != tt.exVal {
						t.Errorf("GetItems() check item has proper status. Expected: %v, got: %v", tt.exVal, item.State)
					}
				}
			}
		})
	}
}

func TestGetItem(t *testing.T) {
	type test struct {
		id   int
		name string
	}
	//var compare test
	tests := []struct {
		name string
		res  test
	}{
		{name: "test0",
			res: test{id: 0,
				name: "light"},
		},
		{name: "test1",
			res: test{id: 1,
				name: "power switch"},
		},
	}
	/* 	{
		{"fakeSwitch", {0, "light"}},
		{"fakeSwitch", {1, "power switch"}},
	} */

	nhc.SendCommand(testutil.MyCmd.Stringify())
	time.Sleep(300 * time.Millisecond)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Println("starting test ", tt.name)
			got, ok := GetItem(tt.res.id)
			if !ok {
				t.Errorf("test %v failed for item %v", tt.name, tt.res.name)
			}
			fmt.Println("result: ", got.Name)
		})
	}
	/* 		for _, tt := range tests {
	   		t.Run(tt.name, func(t *testing.T) {
	   			got := GetItem(tt.res.id)
	   			compare.res.id = got.res.ID
	   			compare.res.name = got.res.name
	   			if got != compare {
	   				t.Errorf("test %v failed for item %v", tt.name, tt.res.name)
	   			}
	   		}
	   	} */
}
