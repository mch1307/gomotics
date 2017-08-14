package nhc

import (
	"encoding/json"
	"fmt"

	"github.com/mch1307/gomotics/db"
	"github.com/mch1307/gomotics/log"
	"github.com/mch1307/gomotics/types"
)

var (
	nhcActions   []types.NhcAction
	nhcLocations []types.NhcLocation
	nhcEvent     []types.NhcEvent
)

// Route parse and route incoming message the right handler
func Route(msg types.NhcMessage) {
	if msg.Cmd == "listlocations" {
		err := json.Unmarshal(msg.Data, &nhcLocations)
		if err != nil {
			fmt.Println("error unmarshalling listlocations")
			log.Warn(err)
			panic(err)
		}
		for idx := range nhcLocations {
			db.SaveNhcLocation(nhcLocations[idx])
		}
	} else if msg.Cmd == "listactions" {
		err := json.Unmarshal(msg.Data, &nhcActions)
		if err != nil {
			fmt.Println("error unmarshalling listactions")
			log.Warn(err)
			panic(err)
		}
		for idx := range nhcActions {
			db.SaveNhcAction(nhcActions[idx])
		}
	} else if msg.Event == "listactions" {
		err := json.Unmarshal(msg.Data, &nhcEvent)
		if err != nil {
			fmt.Println("error unmarshalling listactions event")
			log.Warn(err)
			panic(err)
		}
		for idx := range nhcEvent {
			db.ProcessNhcEvent(nhcEvent[idx])
		}

	}
}
