package server

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/mch1307/gomotics/config"
	"github.com/mch1307/gomotics/db"
	"github.com/mch1307/gomotics/log"
	"github.com/mch1307/gomotics/types"
)

var apiURL string

func initJeedomLocations() {
	locs := GetJeedomObjects()
	for _, val := range locs {
		db.SaveJeedomLocation(val)
	}
}

func initJeedomEquipments() {
	eqs := GetJeedomEquipments()
	for _, eq := range eqs {
		db.SaveJeedomItem(eq)
		cmds := GetJeedomCMDs(eq.ID)
		for _, cmd := range cmds {
			db.SaveJeedomCMD(cmd)
		}
	}
}

// JeedomInit Initialize jeedom internal "db"
func JeedomInit() {
	apiURL = "http://" + strings.Replace(config.Conf.ServerConfig.GMHostPort, "\"", "", -1) + "/api/v1/jeedom"
	log.Debug(apiURL)
	scriptInstalled, err := isJeedomPluginInstalled("script")
	if err != nil {
		log.Warn("Unable to acces Jeedom: ", err)
		config.Conf.JeedomConfig.Enabled = false
		return
	}
	if !isJeedomVersionOK() {
		log.Warn("Non compliant Jeedom version")
		config.Conf.JeedomConfig.Enabled = false
		return
	}
	// install if not installed??
	if scriptInstalled {
		initJeedomLocations()
		initJeedomEquipments()
		db.FillNHCItems()
		if config.Conf.JeedomConfig.AutoCreateObjects {
			// check which locations are missing. If any, we create them in Jeedom
			// and re-match with NHC locations
			missingLocations := db.GetMissingJeedomObjects()
			if len(missingLocations) > 0 {
				log.Debug("missing locations:", missingLocations)
				for _, n := range missingLocations {
					_ = CreateJeedomObject(n)
				}
				initJeedomLocations()
				db.FillNHCItems()
			}
			// check which equipments are missing in Jeedom. If any, we create them in
			// Jeedom with relevant cmds and re-match with NHC actions
			missingEquipments := db.GetMissingJeedomEquipment()
			log.Debug("missing equipments:", missingEquipments)
			for _, eq := range missingEquipments {
				createJeedomEquipment(eq)
			}
			initJeedomEquipments()
			db.FillNHCItems()
		}
		syncJeedomItemsState()
	} else {
		log.Warn("Script Plugin is not installed, disabling Jeedom")
		config.Conf.JeedomConfig.Enabled = false
	}
}

func createJeedomEquipment(eq types.NHCItem) {
	if eq.Type == "dimmer" || eq.Type == "switch" {
		newJeedomEq, err := CreateJeedomEquipment(db.GetJeedomLocationID(eq.Location), eq.Name)
		if err != nil {
			log.Warn(err)
		}
		if len(newJeedomEq) > 0 {
			var stateID string
			if eq.Type == "switch" {
				tempCMD := makeJeedomCMDFromTemplate(eq.Type, "state")
				if jeeCMD, ok := tempCMD.(types.JeedomSwitchStateCMD); ok {
					jeeCMD.Apikey = config.Conf.JeedomConfig.APIKey
					jeeCMD.EqLogicID = newJeedomEq
					stateID, _ = CreateJeedomCMD(jeeCMD)
					log.Debug("##################### id:", stateID)
				}
				tempCMD = makeJeedomCMDFromTemplate(eq.Type, "on")
				if jeeCMD, ok := tempCMD.(types.JeedomSwitchOnOffCMD); ok {
					jeeCMD.Apikey = config.Conf.JeedomConfig.APIKey
					jeeCMD.Name = "on"
					jeeCMD.EqLogicID = newJeedomEq
					jeeCMD.Value = stateID
					jeeCMD.Configuration.UpdateCmdID = stateID
					jeeCMD.Configuration.Request = apiURL + "/#eqLogic_id#/100"
					jeeCMD.Configuration.UpdateCmdToValue = "100"
					jeeCMD.Order = "1"
					i, _ := CreateJeedomCMD(jeeCMD)
					log.Debug("##################### id:", i)
					jeeCMD.Name = "off"
					jeeCMD.Configuration.Request = apiURL + "/#eqLogic_id#/0"
					jeeCMD.Configuration.UpdateCmdToValue = "0"
					jeeCMD.Order = "2"
					i, _ = CreateJeedomCMD(jeeCMD)
					log.Debug("##################### id:", i)
				}
				tempCMD = makeJeedomCMDFromTemplate(eq.Type, "updState")
				if jeeCMD, ok := tempCMD.(types.JeedomSwitchUpdStateCMD); ok {
					jeeCMD.Apikey = config.Conf.JeedomConfig.APIKey
					jeeCMD.EqLogicID = newJeedomEq
					jeeCMD.Value = stateID
					jeeCMD.Configuration.UpdateCmdID = stateID
					i, _ := CreateJeedomCMD(jeeCMD)
					log.Debug("##################### id:", i)
				}
			} else if eq.Type == "dimmer" {
				tempCMD := makeJeedomCMDFromTemplate(eq.Type, "state")
				if jeeCMD, ok := tempCMD.(types.JeedomDimmerStateCMD); ok {
					jeeCMD.Apikey = config.Conf.JeedomConfig.APIKey
					jeeCMD.EqLogicID = newJeedomEq
					//jeeCMD.Value = stateID
					stateID, _ = CreateJeedomCMD(jeeCMD)
					log.Debug("##################### id:", stateID)
				}
				tempCMD = makeJeedomCMDFromTemplate(eq.Type, "dim")
				if jeeCMD, ok := tempCMD.(types.JeedomDimmerDimCMD); ok {
					jeeCMD.Apikey = config.Conf.JeedomConfig.APIKey
					jeeCMD.EqLogicID = newJeedomEq
					jeeCMD.Value = stateID
					jeeCMD.Configuration.UpdateCmdID = stateID
					jeeCMD.Configuration.Request = apiURL + "/#eqLogic_id#/#slider#"
					i, _ := CreateJeedomCMD(jeeCMD)
					log.Debug("##################### id:", i)
				}
				tempCMD = makeJeedomCMDFromTemplate(eq.Type, "updState")
				if jeeCMD, ok := tempCMD.(types.JeedomDimmerUpdStateCMD); ok {
					jeeCMD.Apikey = config.Conf.JeedomConfig.APIKey
					jeeCMD.EqLogicID = newJeedomEq
					jeeCMD.Value = stateID
					jeeCMD.Configuration.UpdateCmdID = stateID
					i, _ := CreateJeedomCMD(jeeCMD)
					log.Debug("##################### id:", i)
				}
			}
		}
	}
}

func syncJeedomItemsState() {
	items := db.GetNHCItems()
	for _, item := range items {
		if len(item.JeedomID) > 0 {
			UpdateJeedomState(item)
		}
	}

}

func makeJeedomCMDFromTemplate(typ, cla string) interface{} {
	if typ == "switch" {
		switch cla {
		case "state":
			var ret types.JeedomSwitchStateCMD
			json.Unmarshal([]byte(jeedomSwitchStateCMD), &ret)
			log.Debug("state")
			return ret
		case "on", "off":
			log.Debug("onOff")
			var ret types.JeedomSwitchOnOffCMD
			json.Unmarshal([]byte(jeedomSwitchOnOffCMD), &ret)
			return ret
		case "updState":
			var ret types.JeedomSwitchUpdStateCMD
			json.Unmarshal([]byte(jeedomSwitchUpdStateCMD), &ret)
			log.Debug("updState")
			return ret
		}
	} else if typ == "dimmer" {
		switch cla {
		case "state":
			var ret types.JeedomDimmerStateCMD
			json.Unmarshal([]byte(jeedomDimmerStateCMD), &ret)
			log.Debug("state")
			return ret
		case "dim":
			var ret types.JeedomDimmerDimCMD
			json.Unmarshal([]byte(jeedomDimmerDimCMD), &ret)
			log.Debug("dim")
			return ret
		case "updState":
			var ret types.JeedomDimmerUpdStateCMD
			json.Unmarshal([]byte(jeedomDimmerUpdStateCMD), &ret)
			log.Debug("updState")
			return ret
		}
	}
	return nil
}

// UpdateJeedomState updates the device's status in Jeedom
func UpdateJeedomState(item types.NHCItem) error {

	var err error
	curItem, _ := db.GetNHCItem(item.ID)
	if curItem.JeedomState != strconv.Itoa(item.State) {

		cli := http.Client{Timeout: time.Second * 10}
		log.Debug("updjeedom: ", item)
		req, _ := http.NewRequest(http.MethodGet, config.Conf.JeedomConfig.URL, nil)
		qry := req.URL.Query()
		qry.Add("apikey", config.Conf.JeedomConfig.APIKey)
		qry.Add("type", "cmd")
		qry.Add("id", item.JeedomUpdState)
		qry.Add(item.JeedomSubType, strconv.Itoa(item.State))
		req.URL.RawQuery = qry.Encode()
		log.Debug("jeedom upd url: ", req.URL.String())
		_, err = cli.Do(req)
		if err != nil {
			log.Warn(err)
		}
	}
	return err
}

// makeRPCArgs returns prepared Jeedom RPC basic args
func makeRPCArgs() (rpcArgs types.JsonRpcArgs, rpcParams types.JeedomRPCParams) {
	var args types.JsonRpcArgs
	var params types.JeedomRPCParams
	args.Jsonrpc = "2.0"
	args.ID = "0"
	params.Apikey = config.Conf.JeedomConfig.APIKey
	//args.Params.Apikey = config.Conf.JeedomConfig.APIKey

	return args, params
}

// newJeedomRPCRequest returns a pre-configure http request
func newJeedomRPCRequest(args []byte) *http.Request {
	req, err := http.NewRequest(http.MethodPost, config.Conf.JeedomConfig.URL, bytes.NewBuffer(args))
	if err != nil {
		log.Warn(err)
	}
	req.Header.Set("User-Agent", "gomotics")
	req.Header.Set("Content-Type", "application/json")

	return req
}

// ExecJeedomRPCRequest send JSON RPC request to Jeedom, returns raw result
func ExecJeedomRPCRequest(args *types.JsonRpcArgs) (res []byte, err error) {
	hcli := http.Client{Timeout: time.Second * 10}
	parsedArgs, err := json.Marshal(args)
	if err != nil {
		log.Warn(err)
	}
	log.Debug("sent command: ", args.Method, string(parsedArgs))
	req := newJeedomRPCRequest(parsedArgs)

	resp, err := hcli.Do(req)
	if err != nil {
		log.Warn(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return body, err
}

// CreateJeedomCMD creates a cmd to a Jeedom script equipment
func CreateJeedomCMD(cmd interface{}) (i string, err error) {
	var cmdRsp types.JeedomCreateRsp
	args, _ := makeRPCArgs()
	args.Method = "cmd::save"
	//cmd.Apikey = params.Apikey
	args.Params = cmd
	parsed, _ := json.Marshal(args)
	log.Debug(string(parsed))
	res, err := ExecJeedomRPCRequest(&args)
	if err != nil {
		log.Warn(err)
		return "", err
	}
	err = json.Unmarshal(res, &cmdRsp)
	if err != nil {
		log.Warn(err)
		return "", err
	}
	return cmdRsp.Result.ID, err

}

// CreateJeedomEquipment saves a new script equipment to jeedom
func CreateJeedomEquipment(locID, name string) (i string, err error) {
	var retEquip types.JeedomCreateRsp
	args, params := makeRPCArgs()
	args.Method = "eqLogic::save"
	params.ObjectID = locID
	params.EqTypeName = "script"
	params.IsEnable = "1"
	params.IsVisible = "1"
	params.Name = name
	args.Params = params
	parsedArgs, _ := json.Marshal(args)
	log.Debug("rpc args: ", string(parsedArgs))

	res, err := ExecJeedomRPCRequest(&args)
	if err != nil {
		log.Warn(err)
	}
	log.Debug(string(res))
	jsonErr := json.Unmarshal(res, &retEquip)
	if jsonErr != nil {
		log.Debug("jsonErr: ", jsonErr)
		return "", err
	}

	return retEquip.Result.ID, err

}

// CreateJeedomObject saves the object to jeedom (location)
func CreateJeedomObject(name string) error {
	args, params := makeRPCArgs()
	args.Method = "object::save"
	params.Name = name
	params.Display = jeedomObjectDisplayOpts
	args.Params = params

	_, err := ExecJeedomRPCRequest(&args)
	// TODO: add read of body to check err
	if err != nil {
		return err
	}
	return err
}

func isJeedomVersionOK() bool {
	var jeeVersion types.JeedomVersion
	args, params := makeRPCArgs()
	args.Method = "version"
	args.Params = params
	res, err := ExecJeedomRPCRequest(&args)
	if err != nil {
		return false
	}
	err = json.Unmarshal(res, &jeeVersion)
	if err != nil {
		log.Warn(err)
		log.Debug(string(res))
	}
	vSlice := strings.Split(jeeVersion.Result, ".")
	if len(vSlice) >= 3 {
		major := vSlice[0]
		minor := vSlice[1]
		patch := vSlice[2]
		version, err := strconv.Atoi(major + minor + patch)
		if err != nil {
			log.Debug("error version converting to number:", major+minor+patch)
		}
		if version >= 318 {
			return true
		}
	}
	return false
}

// isJeedomPluginInstalled checks if script plugin is installed
func isJeedomPluginInstalled(p string) (bool, error) {
	var jeedomPlugins types.JeedomPlugins
	var res bool
	args, params := makeRPCArgs()
	args.Method = "plugin::listPlugin"
	args.Params = params

	body, err := ExecJeedomRPCRequest(&args)
	if err != nil {
		log.Warn(err)
	}
	err = json.Unmarshal(body, &jeedomPlugins)
	if err != nil {
		log.Warn(err)
		log.Debug(string(body))
	}
	for _, v := range jeedomPlugins.Result {
		if v.ID == p {
			res = true
		}
	}
	return res, err

}

// GetJeedomObjects gets all objects from Jeedom
func GetJeedomObjects() []types.JeedomLocation {
	var jeedomObjects types.JeedomObjects
	var ret []types.JeedomLocation
	args, params := makeRPCArgs()
	args.Method = "object::all"
	args.Params = params

	body, err := ExecJeedomRPCRequest(&args)
	if err != nil {
		log.Warn(err)
	}
	_ = json.Unmarshal(body, &jeedomObjects)

	for _, val := range jeedomObjects.Result {
		ret = append(ret, val)
	}
	return ret
}

// GetJeedomEquipments list all Jeedom script equipments
func GetJeedomEquipments() []types.JeedomEquipment {
	var jeedomEquipments types.JeedomEquipments
	var ret []types.JeedomEquipment
	args, params := makeRPCArgs()
	args.Method = "eqLogic::byType"
	params.Type = "script"
	args.Params = params

	body, err := ExecJeedomRPCRequest(&args)
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
	args, params := makeRPCArgs()
	args.Method = "cmd::byEqLogicId"
	params.EqLogicID = id
	args.Params = params

	body, err := ExecJeedomRPCRequest(&args)
	if err != nil {
		log.Warn(err)
	}

	_ = json.Unmarshal(body, &jeedomCMDs)
	for _, val := range jeedomCMDs.Result {
		ret = append(ret, val)
	}
	return ret
}
