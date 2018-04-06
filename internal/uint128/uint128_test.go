package uint128

import (
	"bytes"
	"testing"
)

func TestBytes(t *testing.T) {
	buf := [16]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	x := NewFromBytes(buf)

	if xbuf := x.Bytes(); !bytes.Equal(xbuf[:], buf[:]) {
		t.Errorf("unexpected byte slice: got %v, want %v", xbuf, buf)
	}
}

func TestDecimal(t *testing.T) {
	cases := []struct {
		in  string
		ok  bool
		out string
	}{
		{"", false, "0"},
		{"0", true, "0"},
		{"0123456", true, "123456"},
		{"4294967295", true, "4294967295"},
		{"340282366920938463463374607431768211455", true, "340282366920938463463374607431768211455"},
		{"340282366920938463463374607431768211456", false, "0"},
	}

	for _, c := range cases {
		x, err := NewFromString(c.in)
		if (err == nil) != c.ok {
			t.Errorf("unexpected result: got %t, want %t", (err == nil), c.ok)
		}
		if out := x.DecimalString(); out != c.out {
			t.Errorf("unexpected result: got %s, want %s", out, c.out)
		}
	}
}
