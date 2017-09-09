package nhc

import (
	"encoding/json"

	"github.com/mch1307/gomotics/db"

	"github.com/mch1307/gomotics/types"
)

var (
	nhcActions   []types.Action
	nhcLocations []types.Location
	nhcEvent     []types.Event
	nhcInfo      types.NHCSystemInfo
)

// Route parse and route incoming message the right handler
func Route(msg *types.Message) {
	if msg.Cmd == "listlocations" {
		/* Commented err handling as all incoming msg have already been json parsed once */
		/* 		if err := json.Unmarshal(msg.Data, &nhcLocations); err != nil {
			log.Fatal(err)
		} */
		_ = json.Unmarshal(msg.Data, &nhcLocations)
		for idx := range nhcLocations {
			db.SaveLocation(nhcLocations[idx])
			//db.SaveItem(nhcLocations[idx])
		}
	} else if msg.Cmd == "listactions" {
		/* 		if err := json.Unmarshal(msg.Data, &nhcActions); err != nil {
			log.Fatal(err)
		} */
		_ = json.Unmarshal(msg.Data, &nhcActions)
		for idx := range nhcActions {
			db.SaveAction(nhcActions[idx])
			//db.SaveItem(nhcActions[idx])
		}
	} else if msg.Event == "listactions" {
		/* 		if err := json.Unmarshal(msg.Data, &nhcEvent); err != nil {
			log.Errorf("unable to parse message %v, err: %v", msg.Data, err)
		} */
		msg.Event = "dropme"
		_ = json.Unmarshal(msg.Data, &nhcEvent)
		for _, rec := range nhcEvent {
			db.ProcessEvent(rec)
			//db.SaveItem(nhcEvent[idx])
		}
	} else if msg.Cmd == "systeminfo" {
		msg.Cmd = "dropme"
		_ = json.Unmarshal(msg.Data, &nhcInfo)
		db.SaveNhcSysInfo(nhcInfo)
	}
}
