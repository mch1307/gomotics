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
	//nhcConf := config.Conf.NhcConfig

	connectString, err := net.ResolveTCPAddr("tcp", cfg.Host+":"+strconv.Itoa(cfg.Port))
	if err != nil {
		fmt.Println("connNhc ", err)
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
	conn, err := ConnectNhc(&config.Conf.NhcConfig)
	if err != nil {
		log.Errorf("error sending command: %v. Err Msg: %v", cmd, err)
	}
	log.Debug("received command: ", cmd)
	fmt.Fprintf(conn, cmd+"\n")
	return err
}
