package uint128

import "math/big"

// Itoa returns the string representation of x in base 10.
func Itoa(x Int) string {
	return new(big.Int).SetBytes(x.Bytes()).String()
}

// Atoi creates a new Int from s, interpreted in base 10.
func Atoi(s string) (Int, error) {
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
