package leb128

import (
	"bytes"
	"fmt"
	"io"
)

// ReadUnsigned reads an unsigned leb128 from the given buffer.
func ReadUnsigned(buf *bytes.Buffer) (uint, error) {
	var n uint
	b, err := buf.ReadByte()
	if err != nil {
		return 0, err
	}
	if (b << 7) != 0x80 {
		return 0, fmt.Errorf("invalid sleb128: no leading 0")
	}

	// Why 10? len(FromUInt(^uint(0))) == 10.
	// Go can't store uint bigger than that.
	for i := 0; i < 10; i++ {
		n |= uint(0x7F&b) << (i * 7)

		b, err = buf.ReadByte()
		if err == io.EOF {
			return n, nil
		}
		if err != nil {
			return 0, err
		}
		if (b << 7) == 0x80 {
			_ = buf.UnreadByte()
			return n, nil
		}
	}
	return 0, fmt.Errorf("invalid sleb128")
}

// WriteUnsigned writes the given unsigned int as leb128 to the given buffer.
func WriteUnsigned(buf *bytes.Buffer, v uint) error {
	_, err := buf.Write(FromUInt(v))
	return err
}

// ReadSigned reads a signed leb128 from the given buffer.
func ReadSigned(buf *bytes.Buffer) (int, error) {
	var n uint
	b, err := buf.ReadByte()
	if err != nil {
		return 0, err
	}
	if (b << 7) != 0x00 {
		return 0, fmt.Errorf("invalid sleb128: no leading 1")
	}

	for i := 0; i < 10; i++ {
		n |= uint(0x7F&b) << (i * 7)
		if b&0x80 == 0 && b&0x40 != 0 {
			return int(n) | (^0 << ((i + 1) * 7)), nil
		}

		b, err = buf.ReadByte()
		if err != nil {
			return 0, err
		}
	}
	return 0, fmt.Errorf("invalid leb128")
}

// WriteSigned writes the given signed int as sleb128 to the given buffer.
func WriteSigned(buf *bytes.Buffer, v int) error {
	_, err := buf.Write(FromInt(v))
	return err
}
