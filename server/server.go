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

	"github.com/gorilla/mux"
)

const healthMsg = `{"alive":true}`

// Server holds the gomotics app definition
type Server struct {
	Router     *mux.Router
	ListenPort string
	LogLevel   string
	LogFile    string
}

func init() {
	log.Init()
	log.Info("Starting gomotics")
}

// Initialize initialize the server
// also calls the internal in mem db
func (s *Server) Initialize(conf string) {
	config.Initialize(conf)
	s.ListenPort = ":" + strconv.Itoa(config.Conf.ServerConfig.ListenPort)
	s.LogLevel = config.Conf.ServerConfig.LogLevel
	s.LogFile = config.Conf.ServerConfig.LogPath
	s.Router = mux.NewRouter().StrictSlash(true)
	s.intializeRoutes()
	// Initialize NHC in memory db
	nhc.Init(&config.Conf.NhcConfig)
}

func (s *Server) intializeRoutes() {
	s.Router.HandleFunc("/health", Health).Methods("GET")
	s.Router.HandleFunc("/api/v1/nhc/", getNhcItems).Methods("GET")
	s.Router.HandleFunc("/api/v1/nhc/{id}", getNhcItem).Methods("GET")
	s.Router.HandleFunc("/api/v1/nhc/action", nhcCmd).Methods("PUT")

	s.Router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		t, err := route.GetPathTemplate()
		if err != nil {
			return err
		}
		// p will contain regular expression is compatible with regular expression in Perl, Python, and other languages.
		// for instance the regular expression for path '/articles/{id}' will be '^/articles/(?P<v0>[^/]+)$'
		p, err := route.GetPathRegexp()
		if err != nil {
			return err
		}
		m, err := route.GetMethods()
		if err != nil {
			return err
		}
		fmt.Println(strings.Join(m, ","), t, p)
		return nil
	})
}

// Run starts the server process
func (s *Server) Run() {
	go nhc.Listener()
	log.Fatal(http.ListenAndServe(s.ListenPort, s.Router))
}

// Health endpoint for health monitoring
func Health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(healthMsg))
	//fmt.Fprintln(w, "Healthly!")
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
	var myCmd nhc.SimpleCmd
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

func getNhcItem(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	tmp := nhc.GetItems()
	fmt.Println("getnhcItem arg: ", params["idx"])
	var resp nhc.Item
	for _, val := range tmp {
		if i, _ := strconv.Atoi(params["id"]); val.ID == i {
			resp = val
		}
	}
	rsp, _ := json.Marshal(resp)
	w.Header().Set("Content-Type", "application/json")
	w.Write(rsp)
}
