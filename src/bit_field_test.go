package myhdl

import "testing"

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
