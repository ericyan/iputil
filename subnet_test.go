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

func TestSubnetting(t *testing.T) {
	cases := []struct {
		supernet string
		prefix   int
		subnets  []string
	}{
		{"192.168.0.1/24", 26, []string{"192.168.0.0/26", "192.168.0.64/26", "192.168.0.128/26", "192.168.0.192/26"}},
		{"2001:db8::1/48", 50, []string{"2001:db8::/50", "2001:db8:0:4000::/50", "2001:db8:0:8000::/50", "2001:db8:0:c000::/50"}},
	}

	for _, c := range cases {
		_, supernet, _ := net.ParseCIDR(c.supernet)
		subnets := Subnets(supernet, c.prefix)

		for i, subnet := range subnets {
			if subnet.String() != c.subnets[i] {
				t.Errorf("unexpected subnet: got %s, want %s", subnet, c.subnets[i])
			}
		}
	}
}
