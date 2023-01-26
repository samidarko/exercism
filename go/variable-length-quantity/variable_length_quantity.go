package variablelengthquantity

import (
	"sort"
)

func EncodeVarint(input []uint32) (encoding []byte) {
	for _, num := range input {
		encoding = append(encoding, EncodeVLQ(num)...)
	}
	return
}

func DecodeVarint(input []byte) ([]uint32, error) {
	panic("Please implement the DecodeVarint function")
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
