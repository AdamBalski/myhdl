package myhdl

type BitField struct {
	values []byte
	size   int
}

func (bf *BitField) Get(index int) bool {
	byteIndex := index / 8
	bitIndex := index % 8

	return (bf.values[byteIndex]>>bitIndex)&1 == 1
}

func (bf *BitField) Set(index int, value bool) {
	byteIndex := index / 8
	bitIndex := index % 8

	if value {
		bf.values[byteIndex] |= 1 << bitIndex
	} else {
		bf.values[byteIndex] &= ^(1 << bitIndex)
	}
}

func NewBitField(size int) *BitField {
	capacity := size / 8
	if capacity*8 < size {
		capacity++
	}

	return &BitField{make([]byte, capacity), size}
}
