package types

/* import (
	"encoding/json"
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
	Data  json.RawMessage
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
} */

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
