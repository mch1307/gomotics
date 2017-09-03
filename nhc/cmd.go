package nhc

import (
	"fmt"
	"net"
	"strconv"

	"github.com/mch1307/gomotics/config"
	"github.com/mch1307/gomotics/log"
)

// ConnectNhc establish connection to nhc and return the tcp connection
//func ConnectNhc(nhcConf config.NhcConf) (conn *net.TCPConn, err error) {
func ConnectNhc(cfg *config.NhcConf) (conn *net.TCPConn, err error) {

	connectString, err := net.ResolveTCPAddr("tcp", cfg.Host+":"+strconv.Itoa(cfg.Port))
	if err != nil {
		log.Fatal("connNhc ", err)
		return nil, err
	}

	conn, err = net.DialTCP("tcp", nil, connectString)
	if err != nil {
		log.Fatal("error connecting to nhc: ", err)
	}
	//log.Info("Connected to nhc")
	return conn, err
}

// SendCommand send passed command to nhc
func SendCommand(cmd string) error {
	conn, _ := ConnectNhc(&config.Conf.NhcConfig)
	// no error handling as connect will exit in case of issue
	log.Debug("received command: ", cmd)
	fmt.Println("received command: ", cmd)
	fmt.Fprintf(conn, cmd+"\n")
	return nil
}
