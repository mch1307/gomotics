package main

import (
	"time"

	"github.com/mch1307/gomotics/log"
	"github.com/mch1307/gomotics/nhc"
)

func main() {
	log.Info("Starting")
	// Initialize internal "db" with NHC equipments
	// Send list commands to NHC and store the results in memory
	nhc.Init()
	// Startup the NHC listener for getting all events
	go nhc.Listener()
	
	duration := time.Duration(10) * time.Second
	time.Sleep(duration)

	// Start the webserver so server API
	//domo.Start(DomoConfig.ServerConfig.ListenPort)

}
