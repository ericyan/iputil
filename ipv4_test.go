package iputil

import (
	"bytes"
	"testing"
)

func TestParseIPv4(t *testing.T) {
	cases := []struct {
		in  string
		out *IPv4
	}{
		{"invalid", nil},
		{"192.168.0.1", &IPv4{192, 168, 0, 1}},
	}

	for _, c := range cases {
		out := ParseIPv4(c.in)
		if ((out == nil || c.out == nil) && out != c.out) ||
			(out != nil && c.out != nil && !bytes.Equal(out[:], c.out[:])) {
			t.Errorf("unexpected result: got %s, want %s", out, c.out)
		}
	}
}
