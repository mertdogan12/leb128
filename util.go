package leb128

func uint2bit7(n uint) []byte {
	var bs7 []byte
	for n != 0x00 {
		b := byte(n & 0x7F)
		bs7 = append(bs7, b)
		n >>= 7
	}
	return bs7
}

func bytes2bit7(data []byte) []byte {
	var (
		bs  = data
		l   = len(data)
		b7s []byte
		rem byte
		s   byte
	)
	for i := range bs {
		var (
			b  = bs[l-i-1]
			b7 = b << s
		)
		b7 |= rem
		b7 &= 0x7F
		b7s = append(b7s, b7)

		rem = b >> (8 - s - 1)
		if s++; s == 7 {
			b7s = append(b7s, rem)
			s = 0
		}
	}
	return b7s
}
