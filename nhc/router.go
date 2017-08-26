package nhc

import (
	"encoding/json"

	"github.com/mch1307/gomotics/log"
)

var (
	nhcActions   []Action
	nhcLocations []Location
	nhcEvent     []Event
)

// SaveItem save/process nhc item (location, action, event) to in mem "db"
func SaveItem(item MessageIntf) {
	item.Save()
}

// Route parse and route incoming message the right handler
func Route(msg Message) {
	if msg.Cmd == "listlocations" {
		if err := json.Unmarshal(msg.Data, &nhcLocations); err != nil {
			log.Fatal(err)
		}
		for idx := range nhcLocations {
			SaveItem(nhcLocations[idx])
		}
	} else if msg.Cmd == "listactions" {
		if err := json.Unmarshal(msg.Data, &nhcActions); err != nil {
			log.Fatal(err)
		}
		for idx := range nhcActions {
			SaveItem(nhcActions[idx])
		}
	} else if msg.Event == "listactions" {
		if err := json.Unmarshal(msg.Data, &nhcEvent); err != nil {
			log.Fatal(err)
		}
		for idx := range nhcEvent {
			SaveItem(nhcEvent[idx])
		}
	}
}
