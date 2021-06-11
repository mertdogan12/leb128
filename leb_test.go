package leb128_test

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/allusion-be/leb128"
)

func ExampleFromUInt() {
	// MSB ------------------ LSB
	//       10011000011101100101  In raw binary
	//      010011000011101100101  Padded to a multiple of 7 bits
	//  0100110  0001110  1100101  Split into 7-bit groups
	// 00100110 10001110 11100101  Add high 1 bits on all but last (most significant) group to form bytes
	//     0x26     0x8E     0xE5  In hexadecimal
	//
	// -> 0xE5 0x8E 0x26           Output stream (LSB to MSB)
	for _, b := range leb128.FromUInt(624485) {
		fmt.Printf("0x%X: %08b\n", b, b)
	}
	fmt.Println(leb128.LEB128{0xE5, 0x8E, 0x26}.ToUInt())
	// output:
	// 0xE5: 11100101
	// 0x8E: 10001110
	// 0x26: 00100110
	// 624485
}

func ExampleFromInt() {
	// MSB ------------------ LSB
	//          11110001001000000  Binary encoding of 123456
	//      000011110001001000000  As a 21-bit number
	//      111100001110110111111  Negating all bits (one’s complement)
	//      111100001110111000000  Adding one (two’s complement)
	//  1111000  0111011  1000000  Split into 7-bit groups
	// 01111000 10111011 11000000  Add high 1 bits on all but last (most significant) group to form bytes
	//     0x78     0xBB     0xC0  In hexadecimal
	//
	// -> 0xC0 0xBB 0x78           Output stream (LSB to MSB)
	for _, b := range leb128.FromInt(-123456) {
		fmt.Printf("0x%X: %08b\n", b, b)
	}
	fmt.Println(leb128.SLEB128{0xC0, 0xBB, 0x78}.ToInt())
	// output:
	// 0xC0: 11000000
	// 0xBB: 10111011
	// 0x78: 01111000
	// -123456
}

const (
	MaxUint = ^uint(0)
	MinUint = uint(0)
	MaxInt  = int(MaxUint >> 1)
	MinInt  = -MaxInt - 1
)

func TestFromBigUInt(t *testing.T) {
	bi := big.NewInt(624485)
	leb128.FromBigUInt(*bi)
}

func TestFromInt(t *testing.T) {
	for i := MinInt + 1; i < -1; i /= 10 {
		if leb128.FromInt(i).ToInt() != i {
			t.Error(i)
		}
		if leb128.FromInt(-i).ToInt() != -i {
			t.Error(-i)
		}
	}
}

func TestFromUInt(t *testing.T) {
	for i := MinUint + 1; i < MaxUint/10; i *= 10 {
		if leb128.FromUInt(i).ToUInt() != i {
			t.Error(i)
		}
	}
}
