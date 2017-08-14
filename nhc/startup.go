package nhc

import (
	"encoding/json"
	"fmt"

	"github.com/mch1307/gomotics/config"
	"github.com/mch1307/gomotics/db"
	"github.com/mch1307/gomotics/log"
	"github.com/mch1307/gomotics/types"
)

var nhcMessage types.NhcMessage

// Init sends list commands to NHC in order to get all equipments
func Init() {
	fmt.Println("init invoked")
	conn, err := ConnectNhc()
	if err != nil {
		log.Fatal(err)
	}
	reader := json.NewDecoder(conn)
	fmt.Fprintf(conn, config.Conf.NhcConfig.GetLocCmd+"\n")
	if err := reader.Decode(&nhcMessage); err != nil {
		log.Info(err)
		panic(err)
	}
	Route(nhcMessage)
	fmt.Fprintf(conn, config.Conf.NhcConfig.GetEquipCmd+"\n")
	if err := reader.Decode(&nhcMessage); err != nil {
		log.Info(err)
		panic(err)
	}
	Route(nhcMessage)

	// Build the nhc collection
	db.BuildNhcItems()
}
