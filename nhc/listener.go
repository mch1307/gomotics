package nhc

import (
	"encoding/json"
	"fmt"

	"github.com/mch1307/gomotics/config"
	"github.com/mch1307/gomotics/log"
)

// Listener start a connection to nhc host, register itself
// to receive and route all messages from nhc broadcast
func Listener() error {
	var nhcMessage Message

	conn, err := ConnectNhc(&config.Conf.NhcConfig)
	if err != nil {
		fmt.Println("error connecting to nhc")
		return err
	}

	fmt.Fprintf(conn, RegisterCMD+"\n")

	for {
		reader := json.NewDecoder(conn)
		if err := reader.Decode(&nhcMessage); err != nil {
			log.Info(err)
			panic(err)
		}
		if nhcMessage.Cmd == "startevents" {
			log.Info("Listener registered")
			nhcMessage.Cmd = "dropme"
		} else {
			Route(nhcMessage)
		}
	}
}
