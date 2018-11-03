package uint128

import (
	"encoding/binary"
	"errors"
	"math/big"
	"math/bits"
)

var (
	Zero = Int{0x0, 0x0}
	One  = Int{0x0, 0x1}

	ErrOverflow      = errors.New("overflow")
	ErrEmptySlice    = errors.New("empty byte slice")
	ErrInvalidString = errors.New("invalid string")
)

// An Int represents a big-endian unsigned 128-bit integer.
type Int struct {
	Hi, Lo uint64
}

// NewFromBytes creates a new Int from buf, a big-endian byte slice.
func NewFromBytes(buf []byte) (Int, error) {
	if len(buf) == 0 {
		return Zero, ErrEmptySlice
	}

	if len(buf) > 16 {
		return Zero, ErrOverflow
	}

	var b [16]byte
	copy(b[16-len(buf):], buf)

	return Int{
		Hi: binary.BigEndian.Uint64(b[:8]),
		Lo: binary.BigEndian.Uint64(b[8:]),
	}, nil
}

// NewFromString creates a new Int from s, interpreted in base 10.
func NewFromString(s string) (Int, error) {
	i, ok := new(big.Int).SetString(s, 10)
	if !ok {
		return Zero, ErrInvalidString
	}

	// The zero value for an big.Int represents the value 0.
	if len(i.Bytes()) == 0 {
		return Zero, nil
	}

	return NewFromBytes(i.Bytes())
}

// Pow2 returns 2**n, the base-2 exponential of n.
func Pow2(n uint) (Int, error) {
	if n >= 128 {
		return Zero, ErrOverflow
	}

	return Int{0, 1}.Lsh(n), nil
}

// Add returns the sum x+y as a new Int.
func (x Int) Add(y Int) Int {
	lo := x.Lo + y.Lo
	hi := x.Hi + y.Hi
	if x.Lo > lo {
		hi++
	}

	return Int{hi, lo}
}

// Sub returns the difference x-y as a new Int.
func (x Int) Sub(y Int) Int {
	lo := x.Lo - y.Lo
	hi := x.Hi - y.Hi
	if x.Lo < lo {
		hi--
	}

	return Int{hi, lo}
}

// And returns a new Int that is the bitwise AND of two Int values.
func (x Int) And(y Int) Int {
	return Int{x.Hi & y.Hi, x.Lo & y.Lo}
}

// Or returns a new Int that is the bitwise OR of two Int values.
func (x Int) Or(y Int) Int {
	return Int{x.Hi | y.Hi, x.Lo | y.Lo}
}

// Xor returns a new Int that is the bitwise XOR of two Int values.
func (x Int) Xor(y Int) Int {
	return Int{x.Hi ^ y.Hi, x.Lo ^ y.Lo}
}

// Not returns a new Int that is the bitwise NOT of x.
func (x Int) Not() Int {
	return Int{^x.Hi, ^x.Lo}
}

// Lsh moves each bit of x to the left by n bits and returns result as a new Int.
func (x Int) Lsh(n uint) Int {
	if n >= 128 {
		return Zero
	}

	if n >= 64 {
		return Int{x.Lo << (n - 64), 0}
	}

	return Int{(x.Hi << n) | (x.Lo >> (64 - n)), x.Lo << n}
}

// Rsh moves each bit of x to the right by n bits and returns result as a new Int.
func (x Int) Rsh(n uint) Int {
	if n >= 128 {
		return Zero
	}

	if n >= 64 {
		return Int{0, x.Hi >> (n - 64)}
	}

	return Int{x.Hi >> n, (x.Lo >> n) | (x.Hi << (64 - n))}
}

// Len returns the minimum number of bits required to represent x.
func (x Int) BitLen() int {
	if x.Hi == 0 {
		return bits.Len64(x.Lo)
	}

	return 64 + bits.Len64(x.Hi)
}

// LeadingZeros returns the number of leading zero bits in x.
func (x Int) LeadingZeros() int {
	return 128 - x.BitLen()
}

// TrailingZeros returns the number of trailing zero bits in x.
func (x Int) TrailingZeros() int {
	if x.Lo == 0 {
		return 64 + bits.TrailingZeros64(x.Hi)
	}

	return bits.TrailingZeros64(x.Lo)
}

// Cmp compares x and y and returns either -1, 0, or +1 depending on
// whether x is less than, equal to, or greater than y.
func (x Int) Cmp(y Int) int {
	if x.Hi > y.Hi {
		return 1
	}
	if x.Hi < y.Hi {
		return -1
	}

	// Now that x.Hi == y.Hi
	if x.Lo > y.Lo {
		return 1
	}
	if x.Lo < y.Lo {
		return -1
	}

	return 0
}

// IsGreaterThan returns true if x is greater than y.
func (x Int) IsGreaterThan(y Int) bool {
	return x.Cmp(y) == 1
}

// IsLessThan returns true if x is less than y.
func (x Int) IsLessThan(y Int) bool {
	return x.Cmp(y) == -1
}

// IsEqualTo returns true if x is equal to y.
func (x Int) IsEqualTo(y Int) bool {
	if x.Hi == y.Hi && x.Lo == y.Lo {
		return true
	}

	return false
}

// IsEven returns true if x is an even number.
func (x Int) IsEven() bool {
	return x.And(Int{0, 1}) == Zero
}

// IsOdd returns true if x is an odd number.
func (x Int) IsOdd() bool {
	return x.And(Int{0, 1}) != Zero
}

// Bytes returns x as a big-endian byte slice.
func (x Int) Bytes() []byte {
	buf := make([]byte, 16)
	binary.BigEndian.PutUint64(buf[:8], x.Hi)
	binary.BigEndian.PutUint64(buf[8:], x.Lo)
	return buf
}

// String converts x to a string of decimal digits.
func (x Int) String() string {
	return new(big.Int).SetBytes(x.Bytes()).String()
}
