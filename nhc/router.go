package nhc

import "encoding/json"

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
		/* Commented err handling as all incoming msg have already been json parsed once */
		/* 		if err := json.Unmarshal(msg.Data, &nhcLocations); err != nil {
			log.Fatal(err)
		} */
		_ = json.Unmarshal(msg.Data, &nhcLocations)
		for idx := range nhcLocations {
			SaveItem(nhcLocations[idx])
		}
	} else if msg.Cmd == "listactions" {
		/* 		if err := json.Unmarshal(msg.Data, &nhcActions); err != nil {
			log.Fatal(err)
		} */
		_ = json.Unmarshal(msg.Data, &nhcActions)
		for idx := range nhcActions {
			SaveItem(nhcActions[idx])
		}
	} else if msg.Event == "listactions" {
		/* 		if err := json.Unmarshal(msg.Data, &nhcEvent); err != nil {
			log.Errorf("unable to parse message %v, err: %v", msg.Data, err)
		} */
		_ = json.Unmarshal(msg.Data, &nhcEvent)
		for idx := range nhcEvent {
			SaveItem(nhcEvent[idx])
		}
	}
}
