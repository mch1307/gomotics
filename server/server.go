package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/mch1307/gomotics/config"
	"github.com/mch1307/gomotics/log"
	"github.com/mch1307/gomotics/nhc"
	"github.com/mch1307/gomotics/types"

	"github.com/gorilla/mux"
)

// Server holds the gomotics app definition
type Server struct {
	Router     *mux.Router
	ListenPort string
	LogLevel   string
	LogFile    string
}

// Initialize initialize the server
// also calls the internal in mem db
func (s *Server) Initialize() {

	s.ListenPort = ":" + strconv.Itoa(config.Conf.ServerConfig.ListenPort)
	s.LogLevel = config.Conf.ServerConfig.LogLevel
	s.LogFile = config.Conf.ServerConfig.LogPath
	s.Router = mux.NewRouter().StrictSlash(true)
	s.intializeRoutes()
	// Initialize NHC in memory db
	nhc.Init()
}

func (s *Server) intializeRoutes() {
	s.Router.HandleFunc("/health", Health)
	s.Router.HandleFunc("/api/nhc/action", nhcCmd)
	s.Router.HandleFunc("/api/nhc/list", getNhcItems)
}

// Run starts the server process
func (s *Server) Run() {
	go nhc.Listener()
	log.Fatal(http.ListenAndServe(s.ListenPort, s.Router))
}

// Health endpoint for health monitoring
func Health(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Healthly!")
}

// nhcCmd endpoints for sending NHC commands
func nhcCmd(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()
	id, err := strconv.Atoi(strings.Join(vars["id"], ""))
	if err != nil {
		fmt.Println("invalid request: id should be numeric")
	}
	val, err := strconv.Atoi(strings.Join(vars["value"], ""))
	if err != nil {
		fmt.Println("invalid request: value should be numeric")
	}
	var myCmd types.NhcSimpleCmd
	myCmd.Cmd = "executeactions"
	myCmd.ID = id
	myCmd.Value = val
	nhc.SendCommand(myCmd.Stringify())
	fmt.Fprintln(w, "Success")
}

func getNhcItems(w http.ResponseWriter, r *http.Request) {
	tmp := nhc.GetItems()
	resp, _ := json.Marshal(tmp)
	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)
}
