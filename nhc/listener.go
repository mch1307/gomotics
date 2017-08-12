package nhc

import (
	"encoding/json"
	"fmt"

	"github.com/mch1307/go-domo/config"
	"github.com/mch1307/go-domo/log"
)

// Listener start a connection to nhc host, register itself
// to receive and route all messages from nhc broadcast
func Listener() error {
	globalConf, err := config.GetConf()
	if err != nil {
		panic(err)
	}
	nhcConf := globalConf.NhcConfig
	conn, err := ConnectNhc()
	if err != nil {
		return err
	}

	fmt.Fprintf(conn, nhcConf.RegisterCmd+"\n")
	var nhcMessage Message
	var nhcActions []Action
	var nhcLocations []Location

	reader := json.NewDecoder(conn)

	for {
		if err := reader.Decode(&nhcMessage); err != nil {
			log.Warn(err)
			panic(err)
		}
		switch nhcMessage.Cmd {
		case "startevents":
			log.Info("Listener registered")
		case "listactions":
			log.Info("listactions")
			err := json.Unmarshal(nhcMessage.Data, &nhcActions)
			if err != nil {
				log.Warn(err)
				panic(err)
			}
		case "listlocations":
			log.Info("listlocations")
			err := json.Unmarshal(nhcMessage.Data, &nhcLocations)
			if err != nil {
				log.Warn(err)
				panic(err)
			}
		case "event":
			fmt.Println("event received")
		}
	}
}
