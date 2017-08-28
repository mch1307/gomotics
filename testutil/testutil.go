package testutil

import (
	"encoding/json"

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
		json.Unmarshal([]byte(locations), &fakeLocationsMsg)
		nhc.Route(fakeLocationsMsg)
		json.Unmarshal([]byte(actions), &fakeActionsMsg)
		nhc.Route(fakeActionsMsg)
		nhc.BuildItems()
		popFakeRun = true
	}
}
