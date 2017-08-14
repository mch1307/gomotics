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
	/* 	globalConf, err := config.GetConf()
	   	if err != nil {
	   		panic(err)
	   	}
	   	nhcConf := globalConf.NhcConfig */
	nhcConf := config.Conf.NhcConfig

	conn, err := ConnectNhc()
	if err != nil {
		fmt.Println("error connecting to nhc")
		return err
	}

	fmt.Fprintf(conn, nhcConf.RegisterCmd+"\n")
	var nhcMessage Message

	for {
		reader := json.NewDecoder(conn)
		if err := reader.Decode(&nhcMessage); err != nil {
			log.Info(err)
			panic(err)
		}
		fmt.Println("listerner event: ", &nhcMessage.Event)
		fmt.Println("listerner cmd: ", &nhcMessage.Cmd)
		if nhcMessage.Cmd == "startevents" {
			log.Info("Listener registered")
			nhcMessage.Cmd = "dropme"
		} else {
			Route(nhcMessage)
			nhcMessage.Event = "dropme"
		}
	}
}
