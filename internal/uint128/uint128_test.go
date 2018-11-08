package uint128

import (
	"bytes"
	"testing"
)

const maxUint64 = (1<<64 - 1)

func TestComparison(t *testing.T) {
	cases := []struct {
		x   Int
		y   Int
		cmp int
	}{
		{Int{0, 0}, Int{0, 0}, 0},
		{Int{0, 2}, Int{0, 1}, 1},
		{Int{1, 0}, Int{0, 1}, 1},
		{Int{0, 1}, Int{0, maxUint64}, -1},
		{Int{0, maxUint64}, Int{1, 0}, -1},
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

func TestMul(t *testing.T) {
	cases := []struct {
		x string
		y string
		z string // z = x * y
	}{
		{"0", "1", "0"},
		{"1", "18446744073709551616", "18446744073709551616"},
		{"18446744073709551616", "18446744073709551616", "0"},
		{"4294967296", "4294967295", "18446744069414584320"},
		{"9223372036854775808", "9223372036854775809", "85070591730234615875067023894796828672"},
	}

	for _, c := range cases {
		x, _ := NewFromString(c.x)
		y, _ := NewFromString(c.y)
		z, _ := NewFromString(c.z)

		if !x.Mul(y).IsEqualTo(z) {
			t.Errorf("%s * %d != %s, got %s", x, y, z, x.Mul(y))
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
	u1 := Int{14799720563850130797, 11152134164166830811}
	u2 := Int{10868624793753271583, 6542293553298186666}

	expectedAnd := Int{9529907221165552909, 1927615693132931210}
	if !(u1.And(u2)).IsEqualTo(expectedAnd) {
		t.Errorf("unexpected AND result: %s & %s != %s", u1, u2, expectedAnd)
	}

	expectedOr := Int{16138438136437849471, 15766812024332086267}
	if !(u1.Or(u2)).IsEqualTo(expectedOr) {
		t.Errorf("unexpected OR result: %s | %s != %s", u1, u2, expectedOr)
	}

	expectedXor := Int{6608530915272296562, 13839196331199155057}
	if !(u1.Xor(u2)).IsEqualTo(expectedXor) {
		t.Errorf("unexpected XOR result: %s ^ %s != %s", u1, u2, expectedXor)
	}

	expectedNot := Int{maxUint64, maxUint64}
	if !(Zero.Not()).IsEqualTo(expectedNot) {
		t.Errorf("unexpected Not result: ^%s != %s", Zero, expectedNot)
	}

	expectedLsh1 := Int{11152697053990709979, 3857524254624110006}
	if !(u1.Lsh(1)).IsEqualTo(expectedLsh1) {
		t.Errorf("unexpected left shift result: %s >> %d != %s", u1, 1, expectedLsh1)
	}

	expectedLsh64 := Int{11152134164166830811, 0}
	if !(u1.Lsh(64)).IsEqualTo(expectedLsh64) {
		t.Errorf("unexpected left shift result: %s << %d != %s", u1, 64, expectedLsh64)
	}

	expectedRsh1 := Int{7399860281925065398, 14799439118938191213}
	if !(u1.Rsh(1)).IsEqualTo(expectedRsh1) {
		t.Errorf("unexpected right shift result: %s >> %d != %s", u1, 1, expectedRsh1)
	}

	expectedRsh64 := Int{0, 14799720563850130797}
	if !(u1.Rsh(64)).IsEqualTo(expectedRsh64) {
		t.Errorf("unexpected right shift result: %s >> %d != %s", u1, 64, expectedRsh64)
	}

	if !(u1.Lsh(0)).IsEqualTo(u1) || !(u1.Rsh(0)).IsEqualTo(u1) {
		t.Errorf("left/right shift by 0 should be equal to itself")
	}
}

func TestEvenOdd(t *testing.T) {
	cases := []struct {
		x    Int
		even bool
		odd  bool
	}{
		{Int{0, 0}, true, false},
		{Int{0, 1}, false, true},
		{Int{0, 2}, true, false},
		{Int{1, 0}, true, false},
		{Int{1, 1}, false, true},
		{Int{1, 2}, true, false},
	}

	for _, c := range cases {
		if c.x.IsEven() != c.even || c.x.IsOdd() != c.odd {
			t.Errorf("unexpected oddness for %s", c.x)
		}
	}
}

func TestBitCounting(t *testing.T) {
	cases := []struct {
		n             Int
		bitLen        int
		leadingZeros  int
		trailingZeros int
	}{
		{Zero, 0, 128, 128},
		{Int{0, 1984}, 11, 117, 6},
		{Int{1, 1984}, 65, 63, 6},
		{Int{1984, 0}, 75, 53, 70},
		{Int{1984, 1}, 75, 53, 0},
	}

	for _, c := range cases {
		if bitLen := c.n.BitLen(); bitLen != c.bitLen {
			t.Errorf("unexpected bit len for %s: got %d, want %d", c.n, bitLen, c.bitLen)
		}

		if leadingZeros := c.n.LeadingZeros(); leadingZeros != c.leadingZeros {
			t.Errorf("unexpected leading zeros for %s: got %d, want %d", c.n, leadingZeros, c.leadingZeros)
		}

		if trailingZeros := c.n.TrailingZeros(); trailingZeros != c.trailingZeros {
			t.Errorf("unexpected trailing zeros for %s: got %d, want %d", c.n, trailingZeros, c.trailingZeros)
		}
	}
}
