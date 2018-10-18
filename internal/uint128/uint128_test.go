package uint128

import (
	"bytes"
	"testing"
)

const maxUint64 = (1<<64 - 1)

func TestComparison(t *testing.T) {
	cases := []struct {
		x   Uint128
		y   Uint128
		cmp int
	}{
		{Uint128{0, 0}, Uint128{0, 0}, 0},
		{Uint128{0, 2}, Uint128{0, 1}, 1},
		{Uint128{1, 0}, Uint128{0, 1}, 1},
		{Uint128{0, 1}, Uint128{0, maxUint64}, -1},
		{Uint128{0, maxUint64}, Uint128{1, 0}, -1},
	}

	for _, c := range cases {
		if c.x.Cmp(c.y) != c.cmp {
			t.Errorf("unexpected cmp(%s, %s): want %d, got %d", c.x, c.y, c.cmp, c.x.Cmp(c.y))
		}

		switch c.cmp {
		case -1:
			if c.x.IsGreaterThan(c.y) || c.x.IsEqualTo(c.y) || !c.x.IsLessThan(c.y) {
				t.Errorf("helper method is not consistent when cmp == -1")
			}
		case 0:
			if c.x.IsGreaterThan(c.y) || !c.x.IsEqualTo(c.y) || c.x.IsLessThan(c.y) {
				t.Errorf("helper method is not consistent when cmp == 0")
			}
		case 1:
			if !c.x.IsGreaterThan(c.y) || c.x.IsEqualTo(c.y) || c.x.IsLessThan(c.y) {
				t.Errorf("helper method is not consistent when cmp == 1")
			}
		}
	}
}

func TestAddSub(t *testing.T) {
	cases := []struct {
		x string
		y string
		z string // z = x + y
	}{
		{"0", "1", "1"},
		{"1", "0", "1"},
		{"1", "18446744073709551615", "18446744073709551616"},
		{"18446744073709551615", "1", "18446744073709551616"},
		{"36893488147419103231", "36893488147419103232", "73786976294838206463"},
		{"1", "340282366920938463463374607431768211455", "0"},
		{"340282366920938463463374607431768211455", "1", "0"},
	}

	for _, c := range cases {
		x, _ := NewFromString(c.x)
		y, _ := NewFromString(c.y)
		z, _ := NewFromString(c.z)

		if !x.Add(y).IsEqualTo(z) {
			t.Errorf("%s + %d != %s, got %s", x, y, z, x.Add(y))
		}

		if !z.Sub(y).IsEqualTo(x) {
			t.Errorf("%s - %d != %s, got %s", z, y, x, z.Sub(y))
		}
	}
}

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

func TestPow2(t *testing.T) {
	cases := []struct {
		n   uint
		out string
		err error
	}{
		{0, "1", nil},
		{1, "2", nil},
		{2, "4", nil},
		{63, "9223372036854775808", nil},
		{64, "18446744073709551616", nil},
		{127, "170141183460469231731687303715884105728", nil},
		{128, "0", ErrOverflow},
	}

	for _, c := range cases {
		x, err := Pow2(c.n)
		if err != c.err {
			t.Errorf("unexpected error for %d: got %s, want %s", c.n, err, c.err)
		}
		if out := x.String(); out != c.out {
			t.Errorf("unexpected result: got %s, want %s", out, c.out)
		}
	}
}

func TestBitwise(t *testing.T) {
	u1 := Uint128{14799720563850130797, 11152134164166830811}
	u2 := Uint128{10868624793753271583, 6542293553298186666}

	expectedAnd := Uint128{9529907221165552909, 1927615693132931210}
	if !(u1.And(u2)).IsEqualTo(expectedAnd) {
		t.Errorf("unexpected AND result: %s & %s != %s", u1, u2, expectedAnd)
	}

	expectedOr := Uint128{16138438136437849471, 15766812024332086267}
	if !(u1.Or(u2)).IsEqualTo(expectedOr) {
		t.Errorf("unexpected OR result: %s | %s != %s", u1, u2, expectedOr)
	}

	expectedXor := Uint128{6608530915272296562, 13839196331199155057}
	if !(u1.Xor(u2)).IsEqualTo(expectedXor) {
		t.Errorf("unexpected XOR result: %s ^ %s != %s", u1, u2, expectedXor)
	}
}
