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
		{"213.170.200.0/22", "213.170.200.0", "213.170.203.255"},
		{"2a03:d2c0::/30", "2a03:d2c0::", "2a03:d2c3:ffff:ffff:ffff:ffff:ffff:ffff"},
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
		{"213.170.200.0/22", 24, []string{"213.170.200.0/24", "213.170.201.0/24", "213.170.202.0/24", "213.170.203.0/24"}},
		{"2a03:d2c0::/30", 32, []string{"2a03:d2c0::/32", "2a03:d2c1::/32", "2a03:d2c2::/32", "2a03:d2c3::/32"}},
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
