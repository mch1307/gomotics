package server

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/mch1307/gomotics/config"
	"github.com/mch1307/gomotics/db"
	"github.com/mch1307/gomotics/log"
	"github.com/mch1307/gomotics/types"
)

var args types.JsonRpcArgs

// JeedomInit Initialize jeedom internal "db"
func JeedomInit() {
	scriptOK, err := isJeedomPluginInstalled("script")
	if err != nil {
		log.Warn("Unable to acces Jeedom: ", err)
		config.Conf.JeedomConfig.Enabled = false
		return
	}
	if scriptOK {
		locs := GetJeedomObjects()
		for _, val := range locs {
			db.SaveJeedomLocation(val)
		}
		eqs := GetJeedomEquipments()
		for _, eq := range eqs {
			db.SaveJeedomItem(eq)
			cmds := GetJeedomCMDs(eq.ID)
			for _, cmd := range cmds {
				db.SaveJeedomCMD(cmd)
			}
		}
		db.FillNHCItems()
	} else {
		log.Warn("Script Plugin is not installed, disabling Jeedom")
		config.Conf.JeedomConfig.Enabled = false
	}
}

// UpdateJeedomState updates the device's status in Jeedom
func UpdateJeedomState(item types.NHCItem) error {
	cli := http.Client{Timeout: time.Second * 2}
	log.Debug("updjeedom: ", item)
	req, _ := http.NewRequest(http.MethodGet, config.Conf.JeedomConfig.URL, nil)
	qry := req.URL.Query()
	qry.Add("apikey", config.Conf.JeedomConfig.APIKey)
	qry.Add("type", "cmd")
	qry.Add("id", item.JeedomUpdState)
	qry.Add(item.JeedomSubType, strconv.Itoa(item.State))
	req.URL.RawQuery = qry.Encode()
	log.Debug("jeedom upd url: ", req.URL.String())
	_, err := cli.Do(req)
	if err != nil {
		log.Warn(err)
	}

	return err
}

// makeRPCArgs returns prepared Jeedom RPC basic args
func makeRPCArgs() types.JsonRpcArgs {
	var args types.JsonRpcArgs
	args.Jsonrpc = "2.0"
	args.ID = "0"
	args.Params.Apikey = config.Conf.JeedomConfig.APIKey

	return args
}

// newJeedomRPCRequest returns a pre-configure http request
func newJeedomRPCRequest(args []byte) *http.Request {
	req, err := http.NewRequest(http.MethodPut, config.Conf.JeedomConfig.URL, bytes.NewBuffer(args))
	if err != nil {
		log.Warn(err)
	}
	req.Header.Set("User-Agent", "gomotics")
	req.Header.Set("Content-Type", "application/json")

	return req
}

// JeedomRPCRequest send JSON RPC request to Jeedom, returns raw result
func JeedomRPCRequest(args *types.JsonRpcArgs) (res []byte, err error) {
	hcli := http.Client{Timeout: time.Second * 2}
	parsedArgs, err := json.Marshal(args)
	if err != nil {
		log.Warn(err)
	}
	req := newJeedomRPCRequest(parsedArgs)
	resp, err := hcli.Do(req)
	if err != nil {
		log.Warn(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return body, err
}

// isJeedomPluginInstalled checks if script plugin is installed
func isJeedomPluginInstalled(p string) (bool, error) {
	var jeedomPlugins types.JeedomPlugins
	var res bool
	args := makeRPCArgs()
	args.Method = "plugin::listPlugin"

	body, err := JeedomRPCRequest(&args)
	if err != nil {
		log.Warn(err)
	}
	_ = json.Unmarshal(body, &jeedomPlugins)
	for _, v := range jeedomPlugins.Result {
		if v.Name == "script" {
			res = true
		}
	}
	return res, err

}

// GetJeedomObjects gets all objects from Jeedom
func GetJeedomObjects() []types.JeedomLocation {
	var jeedomObjects types.JeedomObjects
	var ret []types.JeedomLocation
	args := makeRPCArgs()
	args.Method = "object::all"

	body, err := JeedomRPCRequest(&args)
	if err != nil {
		log.Warn(err)
	}
	_ = json.Unmarshal(body, &jeedomObjects)

	for _, val := range jeedomObjects.Result {
		ret = append(ret, val)
	}
	return ret
}

func GetJeedomEquipments() []types.JeedomEquipment {
	var jeedomEquipments types.JeedomEquipments
	var ret []types.JeedomEquipment
	args := makeRPCArgs()
	args.Method = "eqLogic::byType"
	args.Params.Type = "script"

	body, err := JeedomRPCRequest(&args)
	if err != nil {
		log.Warn(err)
	}

	_ = json.Unmarshal(body, &jeedomEquipments)
	for _, val := range jeedomEquipments.Result {
		ret = append(ret, val)
	}
	return ret
}

// GetJeedomCMDs returns the list of cmds of a given equipment
func GetJeedomCMDs(id string) []types.JeedomCMD {
	log.Debug("received id: ", id)
	var jeedomCMDs types.JeedomCMDs
	var ret []types.JeedomCMD
	args := makeRPCArgs()
	args.Method = "cmd::byEqLogicId"
	args.Params.EqLogicID = id

	body, err := JeedomRPCRequest(&args)
	if err != nil {
		log.Warn(err)
	}

	_ = json.Unmarshal(body, &jeedomCMDs)
	for _, val := range jeedomCMDs.Result {
		ret = append(ret, val)
	}
	return ret
}
