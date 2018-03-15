package iputil

import (
	"net"
	"testing"
)

func TestAddressFamily(t *testing.T) {
	cases := []struct {
		ip net.IP
		af uint
		v4 bool
		v6 bool
	}{
		{nil, 0, false, false},
		{net.IP{192, 168, 0, 1}, 4, true, false},
		{net.IP{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xff, 0xff, 192, 168, 0, 1}, 6, true, true},
		{net.IP{0x20, 0x01, 0x0d, 0xb8, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01}, 6, false, true},
	}

	for _, c := range cases {
		if af := AddressFamily(c.ip); af != c.af {
			t.Errorf("unexpected af: got %d, want %d", af, c.af)
		}

		if v4 := IsIPv4(c.ip); v4 != c.v4 {
			t.Errorf("unexpected v4: got %t, want %t", v4, c.v4)
		}

		if v6 := IsIPv6(c.ip); v6 != c.v6 {
			t.Errorf("unexpected v6: got %t, want %t", v6, c.v6)
		}
	}
}

func TestParseDecimal(t *testing.T) {
	cases := []struct {
		s  string
		af uint
		ip net.IP
	}{
		{"", 4, nil},
		{"", 6, nil},
		{"0", 0, nil},
		{"0", 4, net.ParseIP("0.0.0.0")},
		{"0", 6, net.ParseIP("::")},
		{"3232235521", 4, net.ParseIP("192.168.0.1")},
		{"4294967295", 4, net.ParseIP("255.255.255.255")},
		{"42540766411282592856903984951653826561", 6, net.ParseIP("2001:db8::1")},
		{"340282366920938463463374607431768211455", 6, net.ParseIP("ffff:ffff:ffff:ffff:ffff:ffff:ffff:ffff")},
	}

	for _, c := range cases {
		if ip := ParseDecimal(c.s, c.af); !ip.Equal(c.ip) {
			t.Errorf("unexpected result: got %s, want %s", ip, c.ip)
		}
	}
}
