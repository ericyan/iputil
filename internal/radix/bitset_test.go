package radix

import (
	"testing"
)

func TestBitSet(t *testing.T) {
	cases := []struct {
		set []uint
		str string
	}{
		{nil, "[00000000 00000000]"},
		{[]uint{0}, "[10000000 00000000]"},
		{[]uint{8}, "[00000000 10000000]"},
		{[]uint{1, 3, 5, 7, 9, 11, 13, 15}, "[01010101 01010101]"},
		{[]uint{0, 2, 4, 6, 8, 10, 12, 14}, "[10101010 10101010]"},
	}

	for _, c := range cases {
		bits := make(bitset, 2)
		for _, idx := range c.set {
			bits.SetTo(idx, 1)
		}

		if bits.String() != c.str {
			t.Errorf("unexpected bits: got %s, want %s", bits, c.str)
		}
	}
}
