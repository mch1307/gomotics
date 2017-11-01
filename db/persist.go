package db

import (
	"encoding/json"
	"strconv"

	"github.com/mch1307/gomotics/log"
	"github.com/mch1307/gomotics/types"
)

var (
	nhcActionsColl   []types.Action
	nhcLocationsColl []types.Location
	nhcItems         []types.NHCItem
	nhcInfo          types.NHCSystemInfo
	jeedomItems      []types.JeedomEquipment
	jeedomLocations  []types.JeedomLocation
	jeedomCMDs       []types.JeedomCMD
	//jeedomObjects []types.JeedomObjectFull
)

// itemType stores the external to internal item types
type itemType struct {
	Provider     string
	ProviderType string
	InternalType string
}

var itemTypes []itemType

func init() {
	itemTypes = []itemType{
		{Provider: "NHC", ProviderType: "1", InternalType: "switch"},
		{Provider: "NHC", ProviderType: "2", InternalType: "dimmer"},
		{Provider: "NHC", ProviderType: "4", InternalType: "blind"},
	}
}

// GetInternalType return the internal device type
func GetInternalType(provider, pType string) (internalType string) {

	for _, item := range itemTypes {
		if item.Provider == provider && item.ProviderType == pType {
			return item.InternalType
		}
	}
	return ""
}

// SaveJeedomLocation save Jeedom object (location) to collection
func SaveJeedomLocation(loc types.JeedomLocation) {
	found := false
	if len(jeedomLocations) > 0 {
		// lookup existing rec
		for idx, val := range jeedomLocations {
			if val.ID == loc.ID {
				jeedomLocations[idx] = loc
				found = true
				log.Debug("Jeedom loc ID %v found and updated", loc.ID)
			}
		}
	}
	if !found {
		jeedomLocations = append(jeedomLocations, loc)
		log.Debug("Jeedom Loc ID %v not found, inserted", loc.ID)
	}
}

// SaveJeedomItem save Jeedom equipment to collection
func SaveJeedomItem(item types.JeedomEquipment) {
	found := false
	// lookup coll if record already exist
	for idx, val := range jeedomItems {
		if val.ID == item.ID {
			jeedomItems[idx] = item
			found = true
			log.Debug("Jeedom Item ID %v found and updated", item.ID)
		}
	}
	if !found {
		jeedomItems = append(jeedomItems, item)
		log.Debug("Jeedom Item ID %v not found, inserted", item.ID)
	}

}

// SaveJeedomCmd save Jeedom equipment to collection
func SaveJeedomCMD(cmd types.JeedomCMD) {
	found := false
	// lookup coll if record already exist
	for idx, val := range jeedomCMDs {
		if val.ID == cmd.ID {
			jeedomCMDs[idx] = cmd
			found = true
			log.Debug("Jeedom CMD ID %v found and updated", cmd.ID)
		}
	}
	if !found {
		jeedomCMDs = append(jeedomCMDs, cmd)
		log.Debug("Jeedom CMD ID %v not found, inserted", cmd.ID)
	}

}

// BuildJeedomItems

// BuildNHCItems builds the collection of NHC items
// "merges" actions and locations
func BuildNHCItems() {
	var nhcItem types.NHCItem
	// loop through NHC raw actions collection
	// and build items collection
	for _, rec := range nhcActionsColl {
		nhcItem.ID = rec.ID
		nhcItem.Name = rec.Name
		nhcItem.Provider = "NHC"
		nhcItem.Type = GetInternalType("NHC", strconv.Itoa(rec.Type))
		nhcItem.State = rec.Value1
		nhcItem.Value2 = rec.Value2
		nhcItem.Value3 = rec.Value3
		tmpLoc := getNHCLocation(rec.Location)
		nhcItem.Location = tmpLoc.Name
		nhcItems = append(nhcItems, nhcItem)
	}
	log.Debug("itemsCollection built")
}

// SaveNHCAction insert/update action in collection
func SaveNHCAction(act types.Action) {
	found := false
	// first lookup if action already exist
	if len(nhcActionsColl) > 0 {
		for idx, item := range nhcActionsColl {
			if item.ID == act.ID {
				nhcActionsColl[idx] = act
				log.Debug("Nhc ID %v found and updated", act.ID)
				found = true
			}
		}
	}
	if !found {
		log.Debugf("Nhc ID %v not found -> inserted", act.ID)
		nhcActionsColl = append(nhcActionsColl, act)
	}
}

// getNHCAction gets nhc action from collection
func getNHCAction(id int) types.Action {
	var ret types.Action
	for idx, val := range nhcActionsColl {
		if nhcActionsColl[idx].ID == id {
			log.Debugf("Nhc ID %v found", id)
			ret = val
		}
	}
	return ret
}

// GetNHCItems lists all NHC items from items collection
func GetNHCItems() []types.NHCItem {
	return nhcItems
}

// GetNHCItem return specific item
func GetNHCItem(id int) (it types.NHCItem, found bool) {
	found = false
	tmp := GetNHCItems()
	var resp types.NHCItem
	for _, val := range tmp {
		if val.ID == id {
			resp = val
			found = true
		}
	}
	return resp, found
}

// SaveNHCLocation insert/update location in collection
func SaveNHCLocation(loc types.Location) {
	// first lookup if action already exist
	found := false
	if len(nhcLocationsColl) > 0 {
		for idx, item := range nhcLocationsColl {
			if item.ID == loc.ID {
				nhcLocationsColl[idx] = loc
				log.Debugf("Nhc location with ID %v found and updated", loc.ID)
				found = true
			}
		}
	}
	if !found {
		nhcLocationsColl = append(nhcLocationsColl, loc)
		log.Debug("Nhc location with ID %v not found -> created", loc.ID)
	}
}

// getNHCLocation gets nhc action from collection
func getNHCLocation(id int) types.Location {
	var ret types.Location
	for idx, val := range nhcLocationsColl {
		if nhcLocationsColl[idx].ID == id {
			log.Debugf("Nhc location with ID %v found", id)
			ret = val
		}
	}
	return ret
}

// ProcessNHCEvent saves new state of nhc equipment to relevant collections
func ProcessNHCEvent(evt types.Event) []byte {
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
	item, found := GetNHCItem(evt.ID)
	var event []byte
	if found {
		event, _ = json.Marshal(item)
		//server.WSPool.Broadcast <- event
	} else {
		log.Debug("no record found: item - ", evt.ID)
	}
	log.Debug("Nhc event processed for NHC action id: ", evt.ID)
	return event
}

// Dump save collections to log file (debug)
func Dump() {
	log.Debug("NHC actions: ", nhcActionsColl)
	log.Debug("NHC actions: ", nhcLocationsColl)
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
