package iputil

import (
	"bytes"
	"testing"
)

func TestBytes(t *testing.T) {
	buf := [16]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	x := new(uint128).setBytes(buf)

	if xbuf := x.bytes(); !bytes.Equal(xbuf[:], buf[:]) {
		t.Errorf("unexpected byte slice: got %v, want %v", xbuf, buf)
	}
}
