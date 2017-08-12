package main

import (
	"fmt"

	"github.com/mch1307/gomotics/config"

	"github.com/mch1307/gomotics/log"
	"github.com/mch1307/gomotics/nhc"
)

func main() {
	log.Info("Starting")
	go nhc.Listener()
	var myCmd nhc.SimpleCmd
	myCmd.Cmd = config.Conf.NhcConfig.GetEquipCmd

	fmt.Println(config.Conf.NhcConfig.GetEquipCmd)
	_ = nhc.SendCommand(myCmd.Stringify())

	// Initialize internal db for storingNHC and Jeedom equipments
	//db.NewStore()
	//rec := map[string]string{"id": "0", "name": "terrasse"}
	//rec["id"] = "0"
	//rec["name"] = "terrasse"
	//db.SaveToStore(rec)

	// Start the webserver so server API
	//domo.Start(DomoConfig.ServerConfig.ListenPort)

}
