package nhc

import (
	"net"
	"reflect"
	"testing"
)

func TestDiscover(t *testing.T) {
	tests := []struct {
		name string
		want net.IP
	}{
		{"no nhc on LAN", nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Discover(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Discover() = %v, want %v", got, tt.want)
			}
		})
	}
}
