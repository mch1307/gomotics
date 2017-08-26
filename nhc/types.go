package nhc

import "encoding/json"

const (
	// RegisterCMD holds NHC startevents
	RegisterCMD = "{\"cmd\":\"startevents\"}"
	// ListActions holds NHC listactions
	ListActions = "{\"cmd\":\"listactions\"}"
	// ListLocations holds NHC listlocations
	ListLocations = "{\"cmd\":\"listlocations\"}"
	// ListEnergies holds NHC listenergy
	ListEnergies = "{\"cmd\":\"listenergy\"}"
	// ListThermostats holds NHC listthermostat
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

// Message generic struct to hold nhc messages
// used to identify the message type before further parsing
type Message struct {
	Cmd   string `json:"cmd"`
	Event string `json:"event"`
	Data  json.RawMessage
}

// SimpleCmd type holding a nhc command
type SimpleCmd struct {
	Cmd   string `json:"cmd"`
	ID    int    `json:"id"`
	Value int    `json:"value1"`
}

// Stringify return the string version of SimpleCmd
func (sc SimpleCmd) Stringify() string {
	tmp, _ := json.Marshal(sc)
	return string(tmp)
}

// MessageIntf NHC Message interface
type MessageIntf interface {
	Save()
}

// Action holds one individual nhc action (equipment)
type Action struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Type     int    `json:"type"`
	Location int    `json:"location"`
	Value1   int    `json:"value1"`
	Value2   int    `json:"value2"`
	Value3   int    `json:"value3"`
}

// Event holds an individual event
type Event struct {
	ID    int `json:"id"`
	Value int `json:"value1"`
}

// Location holds one nhc location
type Location struct {
	ID   int
	Name string
}

// Save process to db
func (loc Location) Save() {
	SaveLocation(loc)
}

// Save process to db
func (act Action) Save() {
	SaveAction(act)
}

// Save process to db
func (evt Event) Save() {
	ProcessEvent(evt)
}

// Item NHC equipment definition
type Item struct {
	Provider string `json:"provider"`
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Location string `json:"location"`
	State    int    `json:"state"`
}
