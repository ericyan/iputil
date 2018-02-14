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

func TestIPv4ToUint32(t *testing.T) {
	cases := []struct {
		in  *IPv4
		out uint32
	}{
		{&IPv4{0, 0, 0, 0}, 0},
		{&IPv4{192, 168, 0, 1}, 3232235521},
		{&IPv4{255, 255, 255, 255}, 4294967295},
	}

	for _, c := range cases {
		if out := c.in.Uint32(); out != c.out {
			t.Errorf("unexpected result: got %d, want %d", out, c.out)
		}
	}
}

func TestUint32ToIPv4(t *testing.T) {
	cases := []struct {
		in  uint32
		out *IPv4
	}{
		{0, &IPv4{0, 0, 0, 0}},
		{3232235521, &IPv4{192, 168, 0, 1}},
		{4294967295, &IPv4{255, 255, 255, 255}},
	}

	for _, c := range cases {
		if out := ParseUint32(c.in); !bytes.Equal(out[:], c.out[:]) {
			t.Errorf("unexpected result: got %s, want %s", out, c.out)
		}
	}
}
