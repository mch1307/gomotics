package db

import (
	"encoding/json"

	"github.com/mch1307/gomotics/log"
	"github.com/mch1307/gomotics/types"
	"github.com/mch1307/gomotics/ws"
)

var (
	actionsColl   []types.Action
	locationsColl []types.Location
	items         []types.Item
	nhcInfo       types.NHCSystemInfo
)

// BuildItems builds the collection of NHC items
// "merges" actions and locations
func BuildItems() {
	var nhcItem types.Item
	// loop through NHC raw actions collection
	// and build items collection
	for _, rec := range actionsColl {
		nhcItem.ID = rec.ID
		nhcItem.Name = rec.Name
		nhcItem.Provider = "NHC"
		nhcItem.State = rec.Value1
		tmpLoc := GetLocation(rec.Location)
		nhcItem.Location = tmpLoc.Name
		items = append(items, nhcItem)
	}
	log.Debug("itemsCollection built")
}

// SaveAction insert/update action in collection
func SaveAction(act types.Action) {
	found := false
	// first lookup if action already exist
	if len(actionsColl) > 0 {
		for idx, item := range actionsColl {
			if item.ID == act.ID {
				actionsColl[idx] = act
				log.Debug("Nhc ID %v found and updated", act.ID)
				found = true
			}
		}
	}
	if !found {
		log.Debugf("Nhc ID %v not found -> inserted", act.ID)
		actionsColl = append(actionsColl, act)
	}
}

// GetAction gets nhc action from collection
func GetAction(id int) types.Action {
	var ret types.Action
	for idx, val := range actionsColl {
		if actionsColl[idx].ID == id {
			log.Debugf("Nhc ID %v found", id)
			ret = val
		}
	}
	return ret
}

// GetItems lists all NHC items from items collection
func GetItems() []types.Item {
	return items
}

// GetItem return specific item
func GetItem(id int) (it types.Item, found bool) {
	found = false
	tmp := GetItems()
	var resp types.Item
	for _, val := range tmp {
		if val.ID == id {
			resp = val
			found = true
		}
	}
	return resp, found
}

// SaveLocation insert/update location in collection
func SaveLocation(loc types.Location) {
	// first lookup if action already exist
	found := false
	if len(locationsColl) > 0 {
		for idx, item := range locationsColl {
			if item.ID == loc.ID {
				locationsColl[idx] = loc
				log.Debugf("Nhc location with ID %v found and updated", loc.ID)
				found = true
			}
		}
	}
	if !found {
		locationsColl = append(locationsColl, loc)
		log.Debug("Nhc location with ID %v not found -> created", loc.ID)
	}
}

// GetLocation gets nhc action from collection
func GetLocation(id int) types.Location {
	var ret types.Location
	for idx, val := range locationsColl {
		if locationsColl[idx].ID == id {
			log.Debugf("Nhc location with ID %v found", id)
			ret = val
		}
	}
	return ret
}

// ProcessEvent saves new state of nhc equipment to relevant collections
func ProcessEvent(evt types.Event) {
	for idx := range actionsColl {
		if actionsColl[idx].ID == evt.ID {
			actionsColl[idx].Value1 = evt.Value
		}
	}
	for idx := range items {
		if items[idx].ID == evt.ID {
			items[idx].State = evt.Value
		}
	}
	item, found := GetItem(evt.ID)

	if found {
		event, _ := json.Marshal(item)
		ws.WSPool.Broadcast <- event
	} else {
		log.Debug("no record found: item ", evt.ID)
	}
	log.Debug("Nhc event processed for NHC action id:", evt.ID)
}

// Dump save collections to log file (debug)
func Dump() {
	log.Debug("NHC actions: ", actionsColl)
	log.Debug("NHC actions: ", locationsColl)
}

// SaveNhcSysInfo saves the NHC system information in mem
func SaveNhcSysInfo(nhcSysInfo types.NHCSystemInfo) {
	nhcInfo = nhcSysInfo
	log.Debug(nhcInfo)
}

// GetNhcSysInfo returns the NHC system information
func GetNhcSysInfo() (nhcSysInfo types.NHCSystemInfo) {
	return nhcInfo
}
