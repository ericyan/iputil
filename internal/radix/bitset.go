package radix

import (
	"fmt"
)

// A bitset is a bit array backed by a big-endian byte slice.
type bitset []byte

// div8 divides n by 8 and returns quotient and remainder. Implemented
// using bitwise operations since 8 is a power of 2.
func div8(n uint) (quo, rem uint) {
	return n >> 3, n & 7
}

// Get returns the bit at index i.
//
// It will panic if the index results in a byte index that exceeds the
// number of bytes held by the bitset.
func (s bitset) Get(i uint) uint8 {
	idx, mod := div8(i)

	return (s[idx] >> (7 - mod)) & 1
}

// SetTo sets the bit at index i to the given value.
//
// It will panic if the index results in a byte index that exceeds the
// number of bytes held by the bitset or the input bit is invalid.
func (s bitset) SetTo(i uint, bit uint8) {
	idx, mod := div8(i)

	switch bit {
	case 0:
		s[idx] &^= 1 << (7 - mod)
	case 1:
		s[idx] |= 1 << (7 - mod)
	default:
		panic("bit must be either 0 or 1")
	}
}

// BitLen returns the number of bits held by the bitset.
func (s bitset) BitLen() int {
	return len(s) << 3
}

// String returns the binary representation of the bitset as a string.
func (s bitset) String() string {
	return fmt.Sprintf("%08b", s)
}
