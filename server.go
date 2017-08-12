package main

import (
	"github.com/mch1307/go-domo/config"
	"github.com/mch1307/go-domo/log"
	"github.com/mch1307/go-domo/nhc"
)

// DomoConfig stores the app configuration
var DomoConfig config.GlobalConfig

func main() {
	log.Info("Starting")
	go nhc.Listener()
	var myCmd nhc.SimpleCmd
	myCmd.Cmd = "executeactions"
	myCmd.ID = 31
	myCmd.Value = 0
	_ = nhc.SendCommand(myCmd)

	// Initialize internal db for storingNHC and Jeedom equipments
	//db.NewStore()
	//rec := map[string]string{"id": "0", "name": "terrasse"}
	//rec["id"] = "0"
	//rec["name"] = "terrasse"
	//db.SaveToStore(rec)

	// Start the webserver so server API
	//domo.Start(DomoConfig.ServerConfig.ListenPort)

}
