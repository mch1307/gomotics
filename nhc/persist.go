package nhc

import (
	"github.com/mch1307/gomotics/log"
)

var (
	actionsColl   []Action
	locationsColl []Location
	items         []Item
)

// BuildItems builds the collection of NHC items
// "merges" actions and locations
func BuildItems() {
	var nhcItem Item
	// loop through NHC raw actions collection
	// and build items collection
	for _, rec := range actionsColl {
		nhcItem.ID = rec.ID
		nhcItem.Name = rec.Name
		nhcItem.Provider = "NHC"
		nhcItem.State = rec.Value1
		tmpLoc := getLocation(rec.Location)
		nhcItem.Location = tmpLoc.Name
		items = append(items, nhcItem)
	}
	log.Debug("itemsCollection built")
}

// NewItem instantiate new NhcItem
func NewItem(provider string, id, state int) Item {
	new := Item{}
	nhcAction := GetAction(id)
	new.Provider = provider
	new.ID = id
	new.Name = nhcAction.Name
	new.State = state
	loc := getLocation(nhcAction.Location)
	new.Location = loc.Name
	return new
}

// SaveAction insert/update action in collection
func SaveAction(act Action) {
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
		actionsColl = append(actionsColl, act)
		log.Debug("Nhc ID %v not found -> inserted", act.ID)
	}
}

// GetAction gets nhc action from collection
func GetAction(id int) Action {
	var ret Action
	for idx, val := range actionsColl {
		if actionsColl[idx].ID == id {
			log.Debug("Nhc ID %v found", id)
			ret = val
		}
	}
	return ret
}

// GetItems lists all NHC items from items collection
func GetItems() []Item {
	return items
}

// SaveLocation insert/update location in collection
func SaveLocation(loc Location) {
	// first lookup if action already exist
	found := false
	if len(locationsColl) > 0 {
		for idx, item := range locationsColl {
			if item.ID == loc.ID {
				locationsColl[idx] = loc
				log.Debug("Nhc location with ID %v found and updated", loc.ID)
				found = true
			}
		}
	}
	if !found {
		locationsColl = append(locationsColl, loc)
		log.Debug("Nhc location with ID %v not found -> created", loc.ID)
	}
}

// getLocation gets nhc action from collection
func getLocation(id int) Location {
	var ret Location
	for idx, val := range locationsColl {
		if locationsColl[idx].ID == id {
			log.Debug("Nhc location with ID %v found", id)
			ret = val
		}
	}
	return ret
}

// ProcessEvent saves new state of nhc equipment to relevant collections
func ProcessEvent(evt Event) {
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
	log.Debug("Nhc event processed for NHC action id:", evt.ID)
}
