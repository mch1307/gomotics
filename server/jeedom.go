package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/mch1307/gomotics/config"
	"github.com/mch1307/gomotics/db"
	"github.com/mch1307/gomotics/log"
	"github.com/mch1307/gomotics/types"
)

//const jeeUrl = "http://jeedom.csnet.me/core/api/jeeApi.php"

var args types.JsonRpcArgs

//TODO: feels like repeating code in the Getxx ..

// JeedomInit Initialize jeedom internal "db"
func JeedomInit() {
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

func makeRPCArgs() types.JsonRpcArgs {
	var args types.JsonRpcArgs
	args.Jsonrpc = "2.0"
	args.ID = "0"
	args.Params.Apikey = config.Conf.JeedomConfig.APIKey

	return args
}

func newJeedomRPCRequest(args []byte) *http.Request {
	req, err := http.NewRequest(http.MethodPut, config.Conf.JeedomConfig.URL, bytes.NewBuffer(args))
	if err != nil {
		log.Warn(err)
	}
	req.Header.Set("User-Agent", "gomotics")
	req.Header.Set("Content-Type", "application/json")

	return req
}

// GetJeedomObjects gets all objects (full) from Jeedom
func GetJeedomObjects() []types.JeedomLocation {
	var jeedomObjects types.JeedomObjects
	var ret []types.JeedomLocation
	hcli := http.Client{Timeout: time.Second * 2}
	args := makeRPCArgs()
	args.Method = "object::all"

	parsedArgs, err := json.Marshal(args)
	if err != nil {
		fmt.Println(err)
	}
	req := newJeedomRPCRequest(parsedArgs)
	resp, err := hcli.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
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
	hcli := http.Client{Timeout: time.Second * 2}
	parsedArgs, _ := json.Marshal(args)
	req := newJeedomRPCRequest(parsedArgs)
	resp, err := hcli.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	_ = json.Unmarshal(body, &jeedomEquipments)
	for _, val := range jeedomEquipments.Result {
		ret = append(ret, val)
	}
	return ret
}

func GetJeedomCMDs(id string) []types.JeedomCMD {
	log.Debug("received id: ", id)
	var jeedomCMDs types.JeedomCMDs
	var ret []types.JeedomCMD
	args := makeRPCArgs()
	args.Method = "cmd::byEqLogicId"
	args.Params.EqLogicID = id
	parsedArgs, _ := json.Marshal(args)
	req := newJeedomRPCRequest(parsedArgs)
	hcli := http.Client{Timeout: time.Second * 2}
	resp, err := hcli.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	_ = json.Unmarshal(body, &jeedomCMDs)
	for _, val := range jeedomCMDs.Result {
		ret = append(ret, val)
	}
	return ret
}
