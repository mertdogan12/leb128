package leb128_test

import (
	"bytes"
	"fmt"
	"math/big"
	"testing"

	"github.com/allusion-be/leb128"
)

func TestSigned(t *testing.T) {
	for _, test := range []struct {
		Hex   string
		Value *big.Int
	}{
		{"2A", big.NewInt(42)},
		{"7F", big.NewInt(-1)},
		{"C0BB78", big.NewInt(-123456)},
		{"8089FA00", big.NewInt(2000000)},
		{"808098F4E9B5CAEA00", big.NewInt(60000000000000000)},
		{"EF9BAF8589CF959A92DEB7DE8A929EABB424", newInt(t, "24197857200151252728969465429440056815")},
		{"91E4D0FAF6B0EAE5EDA1C8A1F5EDE1D4CB5B", newInt(t, "-24197857200151252728969465429440056815")},
	} {
		t.Run(test.Hex, func(t *testing.T) {
			e := new(big.Int).Set(test.Value)
			bs, err := leb128.EncodeSigned(e)
			if err != nil {
				t.Fatal(err)
			}
			if h := fmt.Sprintf("%X", bs); h != test.Hex {
				t.Errorf("\n%50s\n%50s", h, test.Hex)
			}

			d := new(big.Int).Set(test.Value)
			bi, err := leb128.DecodeSigned(bytes.NewReader(bs))
			if err != nil {
				t.Fatal(err)
			}
			if bi.Cmp(d) != 0 {
				t.Errorf("%s, \n%s\n%s", test.Hex, d, bi)
			}
		})
	}
}
