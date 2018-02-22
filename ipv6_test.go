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
