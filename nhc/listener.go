package nhc

import (
	"encoding/json"
	"fmt"

	"github.com/mch1307/gomotics/config"
	"github.com/mch1307/gomotics/log"
	"github.com/mch1307/gomotics/types"
)

// Listener start a connection to nhc host, register itself
// to receive and route all messages from nhc broadcast
func Listener() {
	var nhcMessage types.Message

	conn, err := ConnectNhc(&config.Conf.NhcConfig)
	if err != nil {
		log.Fatal("Fatal error connecting to NHC: ", err)
	}

	fmt.Fprintf(conn, RegisterCMD+"\n")

	for {
		reader := json.NewDecoder(conn)
		if err := reader.Decode(&nhcMessage); err != nil {
			log.Errorf("error decoding NHC message %v", err)
		}
		if nhcMessage.Cmd == "startevents" {
			log.Info("listener registered")
			nhcMessage.Cmd = "dropme"
		} else {
			log.Debug("received ", &nhcMessage.Cmd)
			Route(&nhcMessage)
		}
	}
}
