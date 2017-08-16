package main

import (
	"github.com/mch1307/gomotics/log"
	"github.com/mch1307/gomotics/server"
)

func main() {
	log.Info("Starting gomotics")
	// Initialize internal "db" with NHC equipments
	// Send list commands to NHC and store the results in memory
	//nhc.Init()
	// Startup the NHC listener for getting all events
	//go nhc.Listener()
	//go server.RestServer()
	s := server.Server{}
	s.Initialize()
	s.Run()

	// Start the webserver so server API
	//domo.Start(DomoConfig.ServerConfig.ListenPort)

}
