package iputil

import (
	"net"
	"testing"
)

func TestRange(t *testing.T) {
	ipv4Range := &Range{ParseIPv4("192.168.0.100"), ParseIPv4("192.168.0.199")}
	ipv6Range := &Range{ParseIPv6("2001:0db8::1234:1"), ParseIPv6("2001:0db8::5678:1")}

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
