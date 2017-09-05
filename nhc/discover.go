package nhc

import (
	"encoding/hex"
	"fmt"
	"log"
	"net"
	"time"
)

// Discover discover NHC controller by sending UDP pkg on port 10000
// return NHC IP address and boolean
func Discover() net.IP {
	//	var err error
	var nhcConnectString net.IP
	var targetAddr *net.UDPAddr
	data, _ := hex.DecodeString("44")
	addr := net.UDPAddr{IP: net.ParseIP("255.255.255.255"), Port: 10000}

	conn, err := net.ListenUDP("udp4", &net.UDPAddr{IP: net.IPv4zero, Port: 18043})
	if err != nil {
		fmt.Println("err connect: ", err)
	}

	_, err = conn.WriteToUDP(data, &addr)

	b := make([]byte, 1024)
	// goroutine for reading broadcast result
	go func() {
		for {
			_, targetAddr, err = conn.ReadFromUDP(b)
			if err != nil {
				log.Printf("Error: UDP read error: %v", err)
				continue
			}
			// test "nhc" connection to replying IP to make sure targetAddr is a NHC controller
			connectString := net.TCPAddr{IP: targetAddr.IP, Port: 8000}
			if err == nil {
				_, err := net.DialTCP("tcp", nil, &connectString)
				//defer nhConn.Close()
				if err == nil {
					nhcConnectString = connectString.IP
				}
			}
		}
	}()

	time.Sleep(time.Second * 3)
	//defer conn.Close()
	return nhcConnectString
}
