package uint128

import "testing"

func TestStrconv(t *testing.T) {
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
		{"340282366920938463463374607431768211456", "340282366920938463463374607431768211455", ErrOverflow},
	}

	for _, c := range cases {
		x, err := Atoi(c.in)
		if err != c.err {
			t.Errorf("unexpected error for %v: got %s, want %s", c.in, err, c.err)
		}
		if out := Itoa(x); out != c.out {
			t.Errorf("unexpected result: got %s, want %s", out, c.out)
		}
	}
}

func TestItoa(t *testing.T) {
	//x, _ := Atoi("34028236692093846346337460743176821146")
	//println(x.Hi, x.Lo)
}
