package nhc

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/mch1307/gomotics/db"
	"github.com/mch1307/gomotics/types"
)

// NhcCmd endpoints for sending NHC commands
func NhcCmd(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()
	id, err := strconv.Atoi(strings.Join(vars["id"], ""))
	if err != nil {
		fmt.Println("invalid request: id should be numeric")
	}
	val, err := strconv.Atoi(strings.Join(vars["value"], ""))
	if err != nil {
		fmt.Println("invalid request: value should be numeric")
	}
	var myCmd SimpleCmd
	myCmd.Cmd = "executeactions"
	myCmd.ID = id
	myCmd.Value = val
	SendCommand(myCmd.Stringify())
	w.Write([]byte("Success"))
}

// GetNhcItems handler for /api/v1/nhc/
func GetNhcItems(w http.ResponseWriter, r *http.Request) {
	tmp := db.GetItems()
	resp, _ := json.Marshal(tmp)
	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)
}

// GetNhcItem handler for /api/v1/nhc/{id}
func GetNhcItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	found := false
	params := mux.Vars(r)
	tmp := db.GetItems()
	//fmt.Println("getnhcItem arg: ", params["id"])
	var resp types.Item
	for _, val := range tmp {
		if i, _ := strconv.Atoi(params["id"]); val.ID == i {
			fmt.Println("in if", params["id"], i)
			resp = val
			found = true
		}
	}
	if !found {
		fmt.Println("not found")
		//http.Error(w, http.StatusNoContent, "no item matching given id found")
		w.WriteHeader(http.StatusNoContent)
		fmt.Fprint(w, string("no item matching given id found"))
	} else {
		rsp, _ := json.Marshal(resp)
		w.Write(rsp)
	}
}
