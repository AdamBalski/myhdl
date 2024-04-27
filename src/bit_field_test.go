package myhdl

import "testing"

func TestNewBitFieldFromShort(t *testing.T) {
	bitfield := NewBitFieldFromShort(8, 159)
	if bitfield.Get(0) != true {
		t.Errorf("Expected true, got false")
	}
	if bitfield.Get(1) != true {
		t.Errorf("Expected true, got false")
	}
	if bitfield.Get(2) != true {
		t.Errorf("Expected true, got false")
	}
	if bitfield.Get(3) != true {
		t.Errorf("Expected true, got false")
	}
	if bitfield.Get(4) != true {
		t.Errorf("Expected true, got false")
	}
	if bitfield.Get(5) != false {
		t.Errorf("Expected false, got true")
	}
	if bitfield.Get(6) != false {
		t.Errorf("Expected false, got true")
	}
	if bitfield.Get(7) != true {
		t.Errorf("Expected true, got false")
	}
}

func TestNewBitFieldFromShort14Bits(t *testing.T) {
	bitfield := NewBitFieldFromShort(14, 7645)
	number := 0
	for i := range 14 {
		if bitfield.Get(i) {
			number |= 1 << i
		}
	}
	if number != 7645 {
		t.Errorf("Expected 7645, got %d", number)
	}
}

func TestBitField(t *testing.T) {
	bitfield := NewBitField(10)
	bitfield.Set(0, true)
	if bitfield.Get(0) != true {
		t.Errorf("Expected true, got false")
	}
	bitfield.Set(0, false)
	if bitfield.Get(0) != false {
		t.Errorf("Expected false, got true")
	}

	bitfield.Set(9, true)
	if bitfield.Get(9) != true {
		t.Errorf("Expected true, got false")
	}
}

func TestBitField_ToShortOf8Bit(t *testing.T) {
	bitfield := NewBitField(8)
	bitfield.Set(0, true)
	bitfield.Set(1, true)
	bitfield.Set(2, true)
	bitfield.Set(3, true)
	bitfield.Set(4, true)
	bitfield.Set(7, true)
	if bitfield.ToShort() != 159 {
		t.Errorf("Expected 159, got %d", bitfield.ToShort())
	}
}

func TestBitField_ToShort(t *testing.T) {
	bitfield := NewBitField(16)
	bitfield.Set(0, true)
	bitfield.Set(1, true)
	bitfield.Set(2, true)
	bitfield.Set(3, true)
	bitfield.Set(4, true)
	bitfield.Set(5, true)
	bitfield.Set(15, true)
	if bitfield.ToShort() != 32831 {
		t.Errorf("Expected 32831, got %d", bitfield.ToShort())
	}
}

func TestBitField_ToShortIfBitFieldTooLong(t *testing.T) {
	bitfield := NewBitField(17)

	defer func() { _ = recover() }()
	bitfield.ToShort()
	t.Errorf("Expected panic, got no panic")
}
