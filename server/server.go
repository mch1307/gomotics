package server

//TODO: review the http handling (return code, ...)
import (
	"fmt"
	"net"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/mch1307/gomotics/config"
	"github.com/mch1307/gomotics/log"
	//"github.com/mch1307/gomotics/server"
)

// HealthMsg static alive json for health endpoint
const HealthMsg = `{"alive":true}`

//var Clients *ClientPool

// Server holds the gomotics app definition
type Server struct {
	Router     *mux.Router
	ListenPort string
	LogLevel   string
	LogFile    string
}

// Start Initialize and starts the server
func Start(conf string) {
	config.Initialize(conf)
	log.Init()
	s := Server{}
	s.Initialize()
	s.Run()
	NhcInit(&config.Conf.NhcConfig)
	fmt.Println("Starting gomotics")
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
	NhcInit(&config.Conf.NhcConfig)
}

func (s *Server) intializeRoutes() {
	s.Router.HandleFunc("/health", Health).Methods("GET")
	s.Router.HandleFunc("/api/v1/nhc/", GetNhcItems).Methods("GET")
	s.Router.HandleFunc("/api/v1/nhc/{id:[0-9]+}", GetNhcItem).Methods("GET")
	s.Router.HandleFunc("/api/v1/nhc/{id:[0-9]+}/{value:[0-9]+}", NhcCmd).Methods("POST")
	//s.Router.HandleFunc("/api/v1/nhc/action", nhc.NhcCmd).Methods("PUT")
	s.Router.HandleFunc("/api/v1/nhc/info", GetNhcInfo).Methods("GET")
	s.Router.HandleFunc("/events", ServeWebSocket).Methods("GET")

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
		log.Info(strings.Join(m, ","), t, p)
		return nil
	})
}

// Run starts the server process
func (s *Server) Run() {
	go NhcListener()
	log.Fatal(http.ListenAndServe(s.ListenPort, s.Router))
}

// Health endpoint for health monitoring
func Health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(HealthMsg))
}

// GetOutboundIP returns the IP address used for out access
func GetOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	//fmt.Println("get ip: ", localAddr.IP.String())
	return localAddr.IP
}
