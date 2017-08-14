package nhc

import (
	"encoding/json"
	"fmt"

	"github.com/mch1307/gomotics/log"
)

var (
	nhcActions   []Action
	nhcLocations []Location
	nhcEvent     []Event
)

// Route parse and route incoming message the right handler
func Route(msg Message) {

	if msg.Cmd == "listlocations" {
		err := json.Unmarshal(msg.Data, &nhcLocations)
		if err != nil {
			fmt.Println("error unmarshalling listlocations")
			log.Warn(err)
			panic(err)
		}
		fmt.Println("json loc", nhcLocations)
	} else if msg.Cmd == "listactions" {
		err := json.Unmarshal(msg.Data, &nhcActions)
		if err != nil {
			fmt.Println("error unmarshalling listactions")
			log.Warn(err)
			panic(err)
		}
		fmt.Println("json action", nhcActions)
	} else if msg.Event == "listactions" {
		err := json.Unmarshal(msg.Data, &nhcEvent)
		if err != nil {
			fmt.Println("error unmarshalling listactions event")
			log.Warn(err)
			panic(err)
		}
		fmt.Println("json event", nhcEvent)
	}

}
