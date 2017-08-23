package nhc

import (
	"encoding/json"

	"github.com/mch1307/gomotics/log"
	"github.com/mch1307/gomotics/types"
)

var (
	nhcActions   []Action
	nhcLocations []Location
	nhcEvent     []Event
)

// SaveItem save/process nhc item (location, action, event) to in mem "db"
func SaveItem(item Message) {
	item.Save()
}

// Route parse and route incoming message the right handler
// ugly, code repetition with json parsing?
func Route(msg types.NhcMessage) {
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
