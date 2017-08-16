package types

import (
	"encoding/json"
)

const (
	RegisterCMD     = "{\"cmd\":\"startevents\"}"
	ListActions     = "{\"cmd\":\"listactions\"}"
	ListLocations   = "{\"cmd\":\"listlocations\"}"
	ListEnergies    = "{\"cmd\":\"listenergy\"}"
	ListThermostats = "{\"cmd\":\"listthermostat\"}"
)

// GenericItem holds definition of item equipment
type GenericItem struct {
	id        int
	provider  string
	name      string
	location  string
	value     int
	itemType  string
	switchCmd string
}

// NhcMessage generic struct to hold nhc messages
// used to identify the message type before futher parsing
type NhcMessage struct {
	Cmd   string `json:"cmd"`
	Event string `json:"event"`
	//Data []NhcAction `json:"data"`
	//Data []interface{} `json:"data"`
	Data json.RawMessage
}

// NhcAction holds one individual nhc action (equipment)
type NhcAction struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Type     int    `json:"type"`
	Location int    `json:"location"`
	Value1   int    `json:"value1"`
	Value2   int    `json:"value2"`
	Value3   int    `json:"value3"`
}

// NhcEvent holds an individual event
type NhcEvent struct {
	ID    int `json:"id"`
	Value int `json:"value1"`
}

// NhcLocation holds one nhc location
type NhcLocation struct {
	ID   int
	Name string
}

// NhcItem NHC equipment definition
type NhcItem struct {
	Provider string `json:"provider"`
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Location string `json:"location"`
	State    int    `json:"state"`
}

// NhcSimpleCmd type holding a nhc command
type NhcSimpleCmd struct {
	Cmd   string `json:"cmd"`
	ID    int    `json:"id"`
	Value int    `json:"value1"`
}

// Stringify return the string version of SimpleCmd
func (sc NhcSimpleCmd) Stringify() string {
	tmp, _ := json.Marshal(sc)
	return string(tmp)
}

/* // Equipment holds the global equipment definition
type Equipment struct {
	Provider   string
	ProviderID int
	GlobalId   string
	Name       string
	Type       int
	State      int
	LocationID int
}
*/
