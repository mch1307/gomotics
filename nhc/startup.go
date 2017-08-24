package nhc

import (
	"encoding/json"
	"fmt"

	"github.com/mch1307/gomotics/config"
	"github.com/mch1307/gomotics/log"
)

var nhcMessage Message

// Init sends list commands to NHC in order to get all equipments
func Init(cfg *config.NhcConf) {
	conn, err := ConnectNhc(cfg)
	if err != nil {
		log.Fatal(err)
	}
	reader := json.NewDecoder(conn)
	// sends listlocations command to NHC

	fmt.Fprintf(conn, ListLocations+"\n")
	if err := reader.Decode(&nhcMessage); err != nil {
		log.Info(err)
		panic(err)
	}

	Route(nhcMessage)
	fmt.Println("init sending listact")
	// sends listActions command to NHC
	fmt.Fprintf(conn, ListActions+"\n")
	if err := reader.Decode(&nhcMessage); err != nil {
		log.Info(err)
		panic(err)
	}
	Route(nhcMessage)
	defer conn.Close()
	// Build the nhc collection
	BuildItems()
	log.Info("Nhc init done")
}
