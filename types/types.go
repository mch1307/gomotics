package types

import "encoding/json"

//ItemType stores the external to internal item types
type ItemType struct {
	Provider     string `json:"provider,omitempty"`
	ProviderType string `json:"provider_type,omitempty"`
	InternalType string `json:"internal_type,omitempty"`
}

// Message generic struct to hold nhc messages
// used to identify the message type before further parsing
type Message struct {
	Cmd   string `json:"cmd"`
	Event string `json:"event"`
	Data  json.RawMessage
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

// NHCItem represents a registered item
//
// swagger:model
type NHCItem struct {
	// the provider
	// required: true
	Provider string `json:"provider"`
	// the id of this item
	ID int `json:"id"`
	// the type of this item
	// can be switch, dimmer or blind
	Type string `json:"type"`
	// the name of the item
	Name string `json:"name"`
	// the location of the item
	Location string `json:"location"`
	// the current state of the item
	State int `json:"state"`
	// other value of the item
	Value2 int `json:"value2"`
	// other value of the item
	Value3         int `json:"value3"`
	JeedomID       string
	JeedomUpdState string
	JeedomSubType  string
	JeedomState    string
}

// NHCSystemInfo hold the NHC system information
// swagger:model
type NHCSystemInfo struct {
	// NHC Software version
	Swversion       string `json:"swversion"`
	API             string `json:"api"`
	Time            string `json:"time"`
	Language        string `json:"language"`
	Currency        string `json:"currency"`
	Units           int    `json:"units"`
	DST             int    `json:"DST"`
	TZ              int    `json:"TZ"`
	Lastenergyerase string `json:"lastenergyerase"`
	Lastconfig      string `json:"lastconfig"`
}
