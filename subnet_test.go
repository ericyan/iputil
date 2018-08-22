package iputil

import (
	"net"
	"testing"
)

func TestSubnetAddr(t *testing.T) {
	cases := []struct {
		subnet    string
		network   string
		broadcast string
	}{
		{"192.168.0.1/24", "192.168.0.0", "192.168.0.255"},
		{"2001:db8::1/64", "2001:db8::", "2001:db8::ffff:ffff:ffff:ffff"},
	}

	for _, c := range cases {
		_, subnet, _ := net.ParseCIDR(c.subnet)

		if network := NetworkAddr(subnet); !network.Equal(net.ParseIP(c.network)) {
			t.Errorf("unexpected network addr: got %s, want %s", network, c.network)
		}

		if broadcast := BroadcastAddr(subnet); !broadcast.Equal(net.ParseIP(c.broadcast)) {
			t.Errorf("unexpected broadcast addr: got %s, want %s", broadcast, c.broadcast)
		}
	}
}
