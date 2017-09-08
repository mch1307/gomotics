package types

import "encoding/json"

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

// Item NHC equipment definition
type Item struct {
	Provider string `json:"provider"`
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Location string `json:"location"`
	State    int    `json:"state"`
}

// NHCSystemInfo hold the NHC system information
type NHCSystemInfo struct {
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
