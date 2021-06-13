package leb128

import (
	"bytes"
	"fmt"
	"math/big"
)

type LEB128 []byte

var (
	x00 = big.NewInt(0x00)
	x7F = big.NewInt(0x7F)
	x80 = big.NewInt(0x80)
)

func EncodeUnsigned(n *big.Int) (LEB128, error) {
	if n.Cmp(big.NewInt(0)) < 0 {
		return nil, fmt.Errorf("can not leb128 encode negative values")
	}
	for bs := []byte{}; ; {
		i := new(big.Int).And(n, x7F)
		n = n.Div(n, x80)
		if n.Cmp(x00) == 0 {
			b := i.Bytes()
			if len(b) == 0 {
				return []byte{0}, nil
			}
			return append(bs, b...), nil
		} else {
			b := new(big.Int).Or(i, x80)
			bs = append(bs, b.Bytes()...)
		}
	}
}

func DecodeUnsigned(r *bytes.Reader) (*big.Int, error) {
	var (
		weight = big.NewInt(1)
		value  = big.NewInt(0)
	)
	for {
		b, err := r.ReadByte()
		if err != nil {
			return nil, err
		}
		value = value.Add(
			value,
			new(big.Int).Mul(big.NewInt(int64(b&0x7F)), weight),
		)
		weight = weight.Mul(weight, x80)
		if b < 0x80 {
			break
		}
	}
	return value, nil
}
