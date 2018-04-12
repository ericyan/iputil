package uint128

import (
	"bytes"
	"testing"
)

func TestBytes(t *testing.T) {
	cases := []struct {
		in  []byte
		out []byte
		err error
	}{
		{[]byte{}, nil, ErrEmptySlice},
		{[]byte{0, 1, 2, 3, 4, 5, 6, 7}, []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 2, 3, 4, 5, 6, 7}, nil},
		{[]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}, []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}, nil},
	}

	for _, c := range cases {
		x, err := NewFromBytes(c.in)
		if err != c.err {
			t.Errorf("unexpected error for %v: got %s, want %s", c.in, err, c.err)
		}
		if err != nil {
			continue
		}
		if out := x.Bytes(); !bytes.Equal(out, c.out) {
			t.Errorf("unexpected result: got %v, want %v", out, c.out)
		}
	}
}

func TestString(t *testing.T) {
	cases := []struct {
		in  string
		out string
		err error
	}{
		{"", "0", ErrInvalidString},
		{"0", "0", nil},
		{"0123456", "123456", nil},
		{"4294967295", "4294967295", nil},
		{"340282366920938463463374607431768211455", "340282366920938463463374607431768211455", nil},
		{"340282366920938463463374607431768211456", "0", ErrOverflow},
	}

	for _, c := range cases {
		x, err := NewFromString(c.in)
		if err != c.err {
			t.Errorf("unexpected error for %v: got %s, want %s", c.in, err, c.err)
		}
		if out := x.String(); out != c.out {
			t.Errorf("unexpected result: got %s, want %s", out, c.out)
		}
	}
}
