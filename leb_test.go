package leb128_test

import (
	"bytes"
	"fmt"
	"math/big"
	"testing"

	"github.com/allusion-be/leb128"
)

func TestUnsigned(t *testing.T) {
	for _, test := range []struct {
		Hex   string
		Value *big.Int
	}{
		{"00", big.NewInt(0)},
		{"07", big.NewInt(7)},
		{"7F", big.NewInt(127)},
		{"E58E26", big.NewInt(624485)},
		{"80897A", big.NewInt(2000000)},
		{"808098F4E9B5CA6A", big.NewInt(60000000000000000)},
		{"EF9BAF8589CF959A92DEB7DE8A929EABB424", newInt(t, "24197857200151252728969465429440056815")},
	} {
		t.Run(test.Hex, func(t *testing.T) {
			e := new(big.Int)
			e = e.Add(test.Value, e) // Copy
			bs, err := leb128.EncodeUnsigned(e)
			if err != nil {
				t.Fatal(err)
			}
			if h := fmt.Sprintf("%X", bs); h != test.Hex {
				t.Errorf("\n%50s\n%50s", h, test.Hex)
			}

			d := new(big.Int)
			d = d.Add(test.Value, d) // Copy
			bi, err := leb128.DecodeUnsigned(bytes.NewReader(bs))
			if err != nil {
				t.Fatal(err)
			}
			if bi.Cmp(d) != 0 {
				t.Errorf("%s, \n%s\n%s", test.Hex, d, bi)
			}
		})
	}
}

func newInt(t *testing.T, str string) *big.Int {
	bi, ok := new(big.Int).SetString(str, 10)
	if !ok {
		t.Fatal()
	}
	return bi
}
