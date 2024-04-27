package myhdl

type BitField struct {
	values []byte
	size   int
}

func NewBitField(size int) *BitField {
	capacity := size / 8
	if capacity*8 < size {
		capacity++
	}

	return &BitField{make([]byte, capacity), size}
}

func NewBitFieldFromShort(size int, value uint16) *BitField {
	bf := NewBitField(size)
	bf.values[0] = byte(value)
	if size > 8 {
		bf.values[1] = byte(value >> 8)
	}
	return bf
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

func (bf *BitField) ToShort() uint16 {
	if bf.size > 16 {
		panic("Cannot convert more than 16 bits to short")
	}
	var result uint16
	for i := len(bf.values) - 1; i >= 0; i-- {
		result = result<<8 + uint16(bf.values[i])
	}
	return result
}
