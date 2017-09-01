package nhc

import (
	"encoding/json"
)

// TODO: remove all types defined in types package

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
