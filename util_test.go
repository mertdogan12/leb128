package leb128

import (
	"testing"
)

func TestBit7(t *testing.T) {
	t.Run("u624485", func(t *testing.T) {
		var (
			n  = uint(624485)
			bs = []byte{0b1100101, 0b0001110, 0b0100110}
		)
		for i, b := range uint2bit7(n) {
			if b != bs[i] {
				t.Errorf("%d: %07b", i, b)
			}
		}
	})
	t.Run("8x11111111", func(t *testing.T) {
		var (
			n  = repeat(0xFF, 7)
			bs = repeat(0x7F, 8)
		)
		for i, b := range bytes2bit7(n) {
			if b != bs[i] {
				t.Errorf("%d: %07b", i, b)
			}
		}
	})
}

func TestAdd1(t *testing.T) {
	if d := add1([]byte{0b10111111}); d[0] != 0b11000000 {
		t.Errorf("%08b", d)
	}
	if d := add1([]byte{0x01, 0xFF}); d[1] != 0x02 {
		t.Errorf("%08b", d)
	}
	if d := add1(repeat(0xFF, 7)); len(d) != 8 || d[0] != 0x01 {
		t.Errorf("%08b", d)
	}
}

func repeat(v byte, n int) []byte {
	b := make([]byte, n)
	for i := 0; i < n; i++ {
		b[i] = v
	}
	return b
}
