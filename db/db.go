package db

import (
	"fmt"

	"github.com/mch1307/gomotics/types"
)

var (
	nhcActionsColl   []types.NhcAction
	nhcLocationsColl []types.NhcLocation
	nhcItems         []types.NhcItem
)

// Equipment interface for hardware equipment
type Equipment interface {
	Switch() error
	Update() error
}

// BuildNhcItems dd
func BuildNhcItems() {
	var nhcItem types.NhcItem
	// loop through NHC raw actions collection
	for _, rec := range nhcActionsColl {
		nhcItem.ID = rec.ID
		nhcItem.Name = rec.Name
		nhcItem.Provider = "HNC"
		nhcItem.State = rec.Value1
		tmpLoc := getNhcLocation(rec.Location)
		nhcItem.Location = tmpLoc.Name
		fmt.Println("build: ", nhcItem)
	}
}

// NewNhcItem instantiate new NhcItem
func NewNhcItem(provider string, id, state int) types.NhcItem {
	new := types.NhcItem{}
	nhcAction := GetNhcAction(id)
	new.Provider = provider
	new.ID = id
	new.Name = nhcAction.Name
	new.State = state
	loc := getNhcLocation(nhcAction.Location)
	new.Location = loc.Name
	return new
}

// SaveNhcAction insert/update action in collection
func SaveNhcAction(act types.NhcAction) {
	found := false
	// first lookup if action already exist
	if len(nhcActionsColl) > 0 {
		for idx, item := range nhcActionsColl {
			if item.ID == act.ID {
				nhcActionsColl[idx] = act
				found = true
			}
		}
	}
	if !found {
		nhcActionsColl = append(nhcActionsColl, act)
	}
}

// GetNhcAction gets nhc action from collection
func GetNhcAction(id int) types.NhcAction {
	var ret types.NhcAction
	for idx, val := range nhcActionsColl {
		if nhcActionsColl[idx].ID == id {
			ret = val
		}
	}
	return ret
}

// SaveNhcLocation insert/update location in collection
func SaveNhcLocation(loc types.NhcLocation) {
	// first lookup if action already exist
	found := false
	if len(nhcLocationsColl) > 0 {
		fmt.Println("len loc >0")
		for idx, item := range nhcLocationsColl {
			if item.ID == loc.ID {
				nhcLocationsColl[idx] = loc
				found = true
			}
		}
	}
	if !found {
		nhcLocationsColl = append(nhcLocationsColl, loc)
	}
}

// getNhcLocation gets nhc action from collection
func getNhcLocation(id int) types.NhcLocation {
	var ret types.NhcLocation
	for idx, val := range nhcLocationsColl {
		if nhcLocationsColl[idx].ID == id {
			ret = val
		}
	}
	return ret
}

// ProcessNhcEvent saves new state of nhc equipment to relevant collections
func ProcessNhcEvent(evt types.NhcEvent) {
	for idx := range nhcActionsColl {
		if nhcActionsColl[idx].ID == evt.ID {
			nhcActionsColl[idx].Value1 = evt.Value
		}
	}
	for idx := range nhcItems {
		if nhcItems[idx].ID == evt.ID {
			nhcItems[idx].State = evt.Value
		}
	}
}
