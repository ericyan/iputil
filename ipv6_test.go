package iputil

import (
	"bytes"
	"testing"
)

func TestParseIPv6(t *testing.T) {
	cases := []struct {
		in  string
		out *IPv6
	}{
		{"invalid", nil},
		{"::1", &IPv6{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}},
		{"2001:db8::1", &IPv6{32, 1, 13, 184, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}},
	}

	for _, c := range cases {
		out := ParseIPv6(c.in)
		if ((out == nil || c.out == nil) && out != c.out) ||
			(out != nil && c.out != nil && !bytes.Equal(out[:], c.out[:])) {
			t.Errorf("unexpected result: got %s, want %s", out, c.out)
		}
	}
}

func TestIPv6ToBigInt(t *testing.T) {
	cases := []struct {
		in  string
		out string
	}{
		{"::", "0"},
		{"2001:db8::1", "42540766411282592856903984951653826561"},
		{"ffff:ffff:ffff:ffff:ffff:ffff:ffff:ffff", "340282366920938463463374607431768211455"},
	}

	for _, c := range cases {
		if out := ParseIPv6(c.in).BigInt(); out.String() != c.out {
			t.Errorf("unexpected result: got %s, want %s", out, c.out)
		}
	}
}

func TestBigIntToIPv6(t *testing.T) {
	cases := []struct {
		in  string
		out string
	}{
		{"0", "::"},
		{"42540766411282592856903984951653826561", "2001:db8::1"},
		{"340282366920938463463374607431768211455", "ffff:ffff:ffff:ffff:ffff:ffff:ffff:ffff"},
	}

	for _, c := range cases {
		if out := ParseDecimal6(c.in); out.String() != c.out {
			t.Errorf("unexpected result: got %s, want %s", out, c.out)
		}
	}
}
