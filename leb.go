// Package leb128 or Little Endian Base 128 is a form of variable-length code
// compression used to store an arbitrarily large integer in a small number of bytes.
package leb128

import (
	"math/big"
)

type (
	// LEB128 represents an unsigned number encoded using (unsigned) LEB128.
	LEB128 []byte
	// SLEB128 represents a signed number encoded using signed LEB128.
	SLEB128 []byte
)

// FromUInt encodes an unsigned big.Int.
func FromBigUInt(n big.Int) LEB128 {
	return toLEB128(bytes2bit7(n.Bytes()))
}

// FromUInt encodes an unsigned integer.
func FromUInt(n uint) LEB128 {
	return toLEB128(uint2bit7(n))
}

func toLEB128(data []byte) LEB128 {
	for i := range data {
		if i == len(data)-1 {
			break
		}
		data[i] |= 0x80
	}
	return data
}

// FromInt encodes a signed integer.
func FromInt(n int) SLEB128 {
	leb := make([]byte, 0)
	for {
		var (
			b    = byte(n & 0x7F)
			sign = byte(n & 0x40)
		)
		if n >>= 7; sign == 0 && n != 0 || n != -1 && (n != 0 || sign != 0) {
			b |= 0x80
		}
		leb = append(leb, b)
		if b&0x80 == 0 {
			break
		}
	}
	return leb
}

// ToUInt converts the byte slice back the an unsigned integer.
func (l LEB128) ToUInt() uint {
	var n uint
	for i := 0; i < len(l); i++ {
		b := uint(0x7F & l[i])
		n |= b << (i * 7)
	}
	return n
}

// ToInt converts the byte slice back the a signed integer.
func (l SLEB128) ToInt() int {
	var n uint
	for i := 0; i < len(l); i++ {
		b := uint(0x7F & l[i])
		n |= b << (i * 7)
		if b := l[i]; b&0x80 == 0 && b&0x40 != 0 {
			return int(n) | (^0 << ((i + 1) * 7))
		}
	}
	return int(n)
}
