package variablelengthquantity

import (
	"errors"
	"sort"
)

func EncodeVarint(input []uint32) (encoding []byte) {
	for _, num := range input {
		encoding = append(encoding, EncodeVLQ(num)...)
	}
	return
}

func DecodeVarint(input []byte) (decoding []uint32, err error) {
	var num uint32
	isLast := false
	for _, chunk := range input {
		num <<= 7
		isLast = chunk&128 == 0

		if chunk&128 == 128 {
			chunk -= 128
		}
		num += uint32(chunk)

		if isLast {
			decoding = append(decoding, num)
			num = 0
		}
	}
	if !isLast {
		return nil, errors.New("missing last")
	}
	return
}

func EncodeVLQ(num uint32) (encoded []byte) {
	first := true
	for {
		chunk := byte(num & 0x7f)
		num >>= 7
		if first {
			first = false
		} else {
			chunk |= 0x80
		}
		encoded = append(encoded, chunk)
		if num == 0 {
			break
		}
	}

	sort.Slice(encoded, func(a, b int) bool { return a > b })

	return
}
