package testutil

import (
	"encoding/json"
	"fmt"

	"github.com/mch1307/gomotics/nhc"
)

var (
	actions = `{"cmd":"listactions","data":[{"id":0,"name":"light","type":1,"location":1,"value1":0},{"id":1,"name":"power switch","type":1,"location":2,"value1":0}]}
	`
	locations = `{"cmd":"listlocations","data":[{"id":0,"name":""},{"id":1,"name":"Living Room"},{"id":2,"name":"Kitchen"}]}
	`
	actionEvent = `{"event":"listactions","data":[{"id":1,"value1":100}]}
	`
	fakeActionsMsg, fakeLocationsMsg nhc.Message
	popFakeRun                       bool
)

// PopFakeNhc populates in mem db with fake data for UT
func PopFakeNhc() {

	if !popFakeRun {
		//fmt.Println("popFake false")
		if err := json.Unmarshal([]byte(locations), &fakeLocationsMsg); err != nil {
			fmt.Println("Error unmarshalling location")
			panic(err)
		}
		nhc.Route(fakeLocationsMsg)

		if err := json.Unmarshal([]byte(actions), &fakeActionsMsg); err != nil {
			fmt.Println("Error unmarshalling action")
			panic(err)
		}
		nhc.Route(fakeActionsMsg)
		nhc.BuildItems()
		popFakeRun = true
	}
}
