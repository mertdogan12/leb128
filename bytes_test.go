package leb128

import (
	"bytes"
	"testing"
)

func TestReadUnsigned(t *testing.T) {
	var (
		value = uint(624485)
		b     = &bytes.Buffer{}
	)
	for i := 0; i < 3; i++ {
		_ = WriteUnsigned(b, value)
	}

	for b.Len() != 0 {
		v, err := ReadUnsigned(b)
		if err != nil {
			t.Error(err)
			return
		}
		if v != value {
			t.Error(v)
		}
	}
}

func TestReadSigned(t *testing.T) {
	var (
		value = -123456
		b     = &bytes.Buffer{}
	)
	for i := 0; i < 3; i++ {
		_ = WriteSigned(b, value)
	}

	for b.Len() != 0 {
		v, err := ReadSigned(b)
		if err != nil {
			t.Error(err)
			return
		}
		if v != value {
			t.Error(v)
		}
	}
}
