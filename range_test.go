package iputil

import (
	"net"
	"testing"
)

func TestRange(t *testing.T) {
	ipv4Range, _ := NewRange(ParseIPv4("192.168.0.100"), ParseIPv4("192.168.0.199"))
	if str := ipv4Range.String(); str != "192.168.0.100 - 192.168.0.199" {
		t.Errorf("unexpected string for IPv4 range: %s", str)
	}

	ipv6Range, _ := NewRange(ParseIPv6("2001:0db8::1234:1"), ParseIPv6("2001:0db8::5678:1"))
	if str := ipv6Range.String(); str != "2001:db8::1234:1 - 2001:db8::5678:1" {
		t.Errorf("unexpected string for IPv6 range: %s", str)
	}

	if _, err := NewRange(ParseIPv4("192.168.0.100"), ParseIPv6("2001:0db8::5678:1")); err == nil {
		t.Error("error expected for invalid range")
	}
	if _, err := NewRange(ParseIPv4("192.168.0.100"), ParseIPv6("192.168.0.99")); err == nil {
		t.Error("error expected for invalid range")
	}

	cases := []struct {
		Range  *Range
		IP     net.IP
		Result bool
	}{
		{ipv4Range, ParseIPv4("192.168.0.99"), false},
		{ipv4Range, ParseIPv4("192.168.0.100"), true},
		{ipv4Range, ParseIPv4("192.168.0.123"), true},
		{ipv4Range, ParseIPv4("192.168.0.199"), true},
		{ipv4Range, ParseIPv4("192.168.0.200"), false},
		{ipv6Range, ParseIPv6("2001:0db8::1"), false},
		{ipv6Range, ParseIPv6("2001:0db8::1234:1"), true},
		{ipv6Range, ParseIPv6("2001:0db8::3456:1"), true},
		{ipv6Range, ParseIPv6("2001:0db8::5678:1"), true},
		{ipv6Range, ParseIPv6("2001:0db8::abcd:1"), false},
	}

	for _, c := range cases {
		if result := c.Range.Contains(c.IP); result != c.Result {
			t.Errorf("unexpected result for %s: got %t, want %t", c.IP, result, c.Result)
		}
	}
}
