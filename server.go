package main

import (
	"time"

	"github.com/mch1307/gomotics/log"
	"github.com/mch1307/gomotics/nhc"
)

func main() {
	log.Info("Starting")

	//duration := time.Duration(1) * time.Second
	//time.Sleep(duration)
	//var myCmd nhc.SimpleCmd
	//myCmd.Cmd = config.Conf.NhcConfig.GetEquipCmd
	//_ = nhc.SendCommand(nhc.RegisterCMD)
	//fmt.Println(config.Conf.NhcConfig.GetEquipCmd)
	//_ = nhc.SendCommand(nhc.ListActions)
	//_ = nhc.SendCommand(myCmd.Stringify())
	nhc.Init()
	// Initialize internal db for storingNHC and Jeedom equipments
	//db.NewStore()
	//rec := map[string]string{"id": "0", "name": "terrasse"}
	//rec["id"] = "0"
	//rec["name"] = "terrasse"
	//db.SaveToStore(rec)
	go nhc.Listener()

	duration := time.Duration(300) * time.Second
	time.Sleep(duration)

	// Start the webserver so server API
	//domo.Start(DomoConfig.ServerConfig.ListenPort)

}
