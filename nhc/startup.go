package nhc

import (
	"encoding/json"
	"fmt"

	"github.com/mch1307/gomotics/config"
	"github.com/mch1307/gomotics/db"
	"github.com/mch1307/gomotics/log"
	"github.com/mch1307/gomotics/types"
)

const (
	// RegisterCMD holds NHC startevents
	RegisterCMD = "{\"cmd\":\"startevents\"}"
	// ListActions holds NHC listactions
	ListActions = "{\"cmd\":\"listactions\"}"
	// ListLocations holds NHC listlocations
	ListLocations = "{\"cmd\":\"listlocations\"}"
	// ListEnergies holds NHC listenergy
	ListEnergies = "{\"cmd\":\"listenergy\"}"
	// ListThermostats holds NHC listthermostat
	ListThermostats = "{\"cmd\":\"listthermostat\"}"
)

var nhcMessage types.Message

// Init sends list commands to NHC in order to get all equipments
func Init(cfg *config.NhcConf) {
	conn, err := ConnectNhc(cfg)
	if err != nil {
		log.Fatalf("Unable to connect to NHC host: %v. Error: %v", cfg.Host, err)
	}
	reader := json.NewDecoder(conn)

	// sends listlocations command to NHC
	fmt.Fprintf(conn, ListLocations+"\n")
	if err := reader.Decode(&nhcMessage); err != nil {
		log.Fatalf("Unable to parse NHC ListLocations message: %v", err)
	}
	Route(&nhcMessage)

	// sends listActions command to NHC
	fmt.Fprintf(conn, ListActions+"\n")
	if err := reader.Decode(&nhcMessage); err != nil {
		log.Fatalf("Unable to parse NHC ListActions message: %v", err)
	}
	Route(&nhcMessage)

	defer conn.Close()
	// Build the nhc collection
	db.BuildItems()
	if config.Conf.ServerConfig.LogLevel == "DEBUG" {
		db.Dump()
	}
	log.Info("Nhc init done")
}
