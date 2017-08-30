package server

//TODO: review the http handling (return code, ...)
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

const HealthMsg = `{"alive":true}`

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
func (s *Server) Initialize() {
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
	s.Router.HandleFunc("/api/v1/nhc/", GetNhcItems).Methods("GET")
	s.Router.HandleFunc("/api/v1/nhc/{id}", GetNhcItem).Methods("GET")
	s.Router.HandleFunc("/api/v1/nhc/action", NhcCmd).Methods("PUT")

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
	w.Write([]byte(HealthMsg))
	//fmt.Fprintln(w, "Healthly!")
}

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
	var myCmd nhc.SimpleCmd
	myCmd.Cmd = "executeactions"
	myCmd.ID = id
	myCmd.Value = val
	nhc.SendCommand(myCmd.Stringify())
	w.Write([]byte("Success"))
}

// GetNhcItems handler for /api/v1/nhc/
func GetNhcItems(w http.ResponseWriter, r *http.Request) {
	tmp := nhc.GetItems()
	resp, _ := json.Marshal(tmp)
	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)
}

// GetNhcItems handler for /api/v1/nhc/{id}
func GetNhcItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	found := false
	params := mux.Vars(r)
	tmp := nhc.GetItems()
	//fmt.Println("getnhcItem arg: ", params["id"])
	var resp nhc.Item
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
