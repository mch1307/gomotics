package nhc_test

import (
	"fmt"
	"net"
	"reflect"
	"testing"

	"github.com/mch1307/gomotics/log"
	. "github.com/mch1307/gomotics/nhc"
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
		{"stub nhc", getOutboundIP()},
	}
	for _, tt := range tests {
		fmt.Println("starting test ", tt.name)
		if tt.want != nil {
			go stubNHCUDP()
			go stubNHCTCP()
		}
		t.Run(tt.name, func(t *testing.T) {
			if got := Discover(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Discover() = %v, want %v", got, tt.want)
			}
		})
	}
}
