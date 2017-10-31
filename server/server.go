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
	NhcInit(&config.Conf.NhcConfig)
}

func (s *Server) intializeRoutes() {
	s.Router.HandleFunc("/health", Health).Methods("GET")
	// swagger:route GET /nhc List NHC registered items
	//
	// Lists all registered NHC items.
	//
	// This will show all registered NHC items.
	//
	//     Consumes:
	//     - application/json
	//     - application/x-protobuf
	//
	//     Produces:
	//     - application/json
	//     - application/x-protobuf
	//
	//     Schemes: http, https, ws, wss
	//
	//
	//     Responses:
	//       default: genericError
	//       200: someResponse
	//       422: validationError
	s.Router.HandleFunc("/api/v1/nhc/", GetNhcItems).Methods("GET")
	// swagger:route GET /nhc/{id:[0-9]+} Find NHC registered item with specific ID
	//
	// Find NHC registered item with specific ID
	//
	// This will show all details of a given registered NHC item.
	//
	//     Consumes:
	//     - application/json
	//     - application/x-protobuf
	//
	//     Produces:
	//     - application/json
	//     - application/x-protobuf
	//
	//     Schemes: http, https, ws, wss
	//
	//
	//     Responses:
	//       default: genericError
	//       200: someResponse
	//       422: validationError
	s.Router.HandleFunc("/api/v1/nhc/{id:[0-9]+}", GetNhcItem).Methods("GET")
	// swagger:route POST /nhc/{id:[0-9]+}/{value:[0-9]+} Update given NHC id with provided value. on-off -> 0-100, intermediate value for dimmer device.
	//
	//
	//     Produces:
	//     - application/json
	//     - application/x-protobuf
	//
	//     Schemes: http, https, ws, wss
	//
	//
	//     Responses:
	//       default: genericError
	//       200: someResponse
	//       422: validationError
	s.Router.HandleFunc("/api/v1/nhc/{id:[0-9]+}/{value:[0-9]+}", NhcCmd).Methods("POST")
	// swagger:route POST /nhc/info Get NHC controller info
	//
	//
	//     Consumes:
	//     - application/json
	//     - application/x-protobuf
	//
	//     Produces:
	//     - application/json
	//     - application/x-protobuf
	//
	//     Schemes: http, https, ws, wss
	//
	//
	//     Responses:
	//       default: genericError
	//       200: someResponse
	//       422: validationError
	s.Router.HandleFunc("/api/v1/nhc/info", GetNhcInfo).Methods("GET")
	// swagger:route GET /events Websocket endpoint providing updated items in realtime
	//
	// Update on items in realtime
	//
	// Websocket endpoint providing updated items in realtime.
	//
	//
	//
	//
	//
	//     Produces:
	//     - application/json
	//
	//
	//     Schemes: ws, wss
	//
	//
	//     Responses:
	//
	//
	//
	s.Router.HandleFunc("/events", ServeWebSocket).Methods("GET")

	s.Router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		t, _ := route.GetPathTemplate()
		// p will contain regular expression is compatible with regular expression in Perl, Python, and other languages.
		// for instance the regular expression for path '/articles/{id}' will be '^/articles/(?P<v0>[^/]+)$'
		p, _ := route.GetPathRegexp()
		m, _ := route.GetMethods()

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

// Start Initialize and start the server
func Start(conf string) {
	config.Initialize(conf)
	log.Init()
	s := Server{}
	s.Initialize()
	s.Run()
	NhcInit(&config.Conf.NhcConfig)
	fmt.Println("Starting gomotics")
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
		return net.IPv4zero
	}
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	defer conn.Close()
	return localAddr.IP
}
