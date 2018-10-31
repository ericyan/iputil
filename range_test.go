package iputil

import (
	"net"
	"testing"
)

var (
	ipv4Range, _ = NewRange(ParseIPv4("192.168.0.100"), ParseIPv4("192.168.0.199"))
	ipv6Range, _ = NewRange(ParseIPv6("2001:db8::1234:0"), ParseIPv6("2001:db8::5678:0"))
)

func TestRange(t *testing.T) {
	if str := ipv4Range.String(); str != "192.168.0.100 - 192.168.0.199" {
		t.Errorf("unexpected string for IPv4 range: %s", str)
	}

	if str := ipv6Range.String(); str != "2001:db8::1234:0 - 2001:db8::5678:0" {
		t.Errorf("unexpected string for IPv6 range: %s", str)
	}

	if _, err := NewRange(ParseIPv4("192.168.0.100"), ParseIPv6("2001:db8::5678:0")); err == nil {
		t.Error("error expected for invalid range")
	}
	if _, err := NewRange(ParseIPv4("192.168.0.100"), ParseIPv4("192.168.0.99")); err == nil {
		t.Error("error expected for invalid range")
	}
}

func TestRangeContains(t *testing.T) {
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
		{ipv6Range, ParseIPv6("2001:0db8::1234:0"), true},
		{ipv6Range, ParseIPv6("2001:0db8::3456:1"), true},
		{ipv6Range, ParseIPv6("2001:0db8::5678:0"), true},
		{ipv6Range, ParseIPv6("2001:0db8::abcd:1"), false},
	}

	for _, c := range cases {
		if result := c.Range.Contains(c.IP); result != c.Result {
			t.Errorf("unexpected result for %s: got %t, want %t", c.IP, result, c.Result)
		}
	}
}

func TestRangeCIDR(t *testing.T) {
	ipv4CIDRs := []string{
		"192.168.0.100/30",
		"192.168.0.104/29",
		"192.168.0.112/28",
		"192.168.0.128/26",
		"192.168.0.192/29",
	}
	for i, cidr := range ipv4Range.CIDR() {
		if cidr.String() != ipv4CIDRs[i] {
			t.Errorf("unexpected CDIR: got %s, want %s", cidr, ipv4CIDRs[i])
		}
	}

	ipv6CIDRs := []string{
		"2001:db8::1234:0/110",
		"2001:db8::1238:0/109",
		"2001:db8::1240:0/106",
		"2001:db8::1280:0/105",
		"2001:db8::1300:0/104",
		"2001:db8::1400:0/102",
		"2001:db8::1800:0/101",
		"2001:db8::2000:0/99",
		"2001:db8::4000:0/100",
		"2001:db8::5000:0/102",
		"2001:db8::5400:0/103",
		"2001:db8::5600:0/106",
		"2001:db8::5640:0/107",
		"2001:db8::5660:0/108",
		"2001:db8::5670:0/109",
		"2001:db8::5678:0/128",
	}
	for i, cidr := range ipv6Range.CIDR() {
		if cidr.String() != ipv6CIDRs[i] {
			t.Errorf("unexpected CDIR: got %s, want %s", cidr, ipv6CIDRs[i])
		}
	}
}
