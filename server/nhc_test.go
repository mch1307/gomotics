package server_test

import (
	"fmt"
	"net"
	"reflect"
	"testing"
	"time"

	"github.com/mch1307/gomotics/log"
	. "github.com/mch1307/gomotics/server"
	"github.com/mch1307/gomotics/testutil"
)

func stubNHCTCP() {
	// listen to incoming tcp connections
	l, err := net.Listen("tcp", "0.0.0.0:8000")
	if err != nil {
		fmt.Println(err)
	}
	defer l.Close()
	_, err = l.Accept()
	if err != nil {
		fmt.Println(err)
	}
}

func stubNHCUDP() {
	// listen to incoming udp packets
	fmt.Println("starting UDP stub")
	pc, err := net.ListenPacket("udp", "0.0.0.0:10000")
	if err != nil {
		log.Fatal(err)
	}
	defer pc.Close()

	//simple read
	buffer := make([]byte, 1024)
	var addr net.Addr
	_, addr, _ = pc.ReadFrom(buffer)

	//simple write
	pc.WriteTo([]byte("NHC Stub"), addr)
}

func getOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}

func TestDiscover(t *testing.T) {

	tests := []struct {
		name string
		want net.IP
	}{
		{"no nhc on LAN", nil},
		//{"stub nhc", getOutboundIP()},
	}
	portCheckIteration := 0
	for _, tt := range tests {
		fmt.Println("starting test ", tt.name)
		if tt.want != nil {
			go stubNHCUDP()
			go stubNHCTCP()
		}
		t.Run(tt.name, func(t *testing.T) {
		GotoTestPort:
			if testutil.IsTCPPortAvailable(18043) {
				if got := Discover(); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("Discover() = %v, want %v", got, tt.want)
				}
			} else {
				portCheckIteration++
				if portCheckIteration < 21 {
					fmt.Printf("UDP 18043 busy, %v retry", portCheckIteration)
					time.Sleep(time.Millisecond * 500)
					goto GotoTestPort
				} else {
					t.Error("Discover failed to get UDP port 18043, test failed")
				}

			}
		})
	}
}
