package uint128

import "math/big"

// cutoff is the smallest n that n*10 will overlow an Int.
var cutoff = Int{1844674407370955161, 11068046444225730970}

// Itoa returns the string representation of x in base 10.
func Itoa(x Int) string {
	return new(big.Int).SetBytes(x.Bytes()).String()
}

// Atoi creates a new Int from s, interpreted in base 10.
func Atoi(s string) (Int, error) {
	if len(s) == 0 {
		return Zero, ErrInvalidString
	}

	x := Zero
	for _, c := range []byte(s) {
		if c < '0' || c > '9' {
			return Max, ErrInvalidString
		}
		digit := Int{0, uint64(c - '0')}

		if x.IsGreaterThan(cutoff) || x.IsEqualTo(cutoff) {
			return Max, ErrOverflow
		}
		x = x.Mul(Int{0, 10})

		x1 := x.Add(digit)
		if x1.IsLessThan(x) {
			return Max, ErrOverflow
		}
		x = x1
	}

	return x, nil
}
