package nhc

import (
	"fmt"
	"net"
	"strconv"

	"github.com/mch1307/gomotics/config"
	"github.com/mch1307/gomotics/log"
)

// ConnectNhc establish connection to nhc andreturn the tcp connection
//func ConnectNhc(nhcConf config.NhcConf) (conn *net.TCPConn, err error) {
func ConnectNhc(cfg *config.NhcConf) (conn *net.TCPConn, err error) {
	//nhcConf := config.Conf.NhcConfig

	connectString, err := net.ResolveTCPAddr("tcp", cfg.Host+":"+strconv.Itoa(cfg.Port))
	if err != nil {
		println("Issue converting host + port: ", err)
		return nil, err
	}

	conn, err = net.DialTCP("tcp", nil, connectString)
	if err != nil {
		println("error connecting to nhc: ", err)
		panic(err)
	}
	//log.Info("Connected to nhc")
	return conn, err
}

// SendCommand send passed command to nhc
func SendCommand(cmd string) error {
	conn, err := ConnectNhc(&config.Conf.NhcConfig)
	if err != nil {
		panic(err)
	}
	log.Debug("received command: ", cmd)
	fmt.Fprintf(conn, cmd+"\n")
	return err
}
