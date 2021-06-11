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

func add1(data []byte) []byte {
	for i := range data {
		j := len(data) - i - 1
		if data[j] == 0xFF {
			data[j] = 0x00
			if i == 0 {
				data = append([]byte{0x01}, data...)
			}
		} else {
			data[j] += 0x01
			break
		}
	}
	return data
}

func add1high(data []byte) []byte {
	for i := range data {
		if i == len(data)-1 {
			break
		}
		data[i] |= 0x80
	}
	return data
}
