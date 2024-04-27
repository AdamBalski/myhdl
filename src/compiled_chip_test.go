package myhdl

import "testing"

func createXor() *CompiledChip {
	return &CompiledChip{
		name:         "Xor",
		inputsNames:  []string{"a", "b"},
		outputsNames: []string{"out"},
		inputs:       NewBitField(2),
		outputs:      NewBitField(1),
		table: []*BitField{
			NewBitFieldFromShort(1, 0),
			NewBitFieldFromShort(1, 1),
			NewBitFieldFromShort(1, 1),
			NewBitFieldFromShort(1, 0),
		},
	}
}

func TestCompiledChip_GetInputIdx(t *testing.T) {
	xor := createXor()
	if xor.GetInputIdx("a") != 0 {
		t.Errorf("Expected 0, got %d", xor.GetInputIdx("a"))
	}
	if xor.GetInputIdx("b") != 1 {
		t.Errorf("Expected 1, got %d", xor.GetInputIdx("b"))
	}
	if xor.GetInputIdx("c") != -1 {
		t.Errorf("Expected -1, got %d", xor.GetInputIdx("c"))
	}
}

func TestCompiledChip_GetOutputIdx(t *testing.T) {
	xor := createXor()
	if xor.GetOutputIdx("out") != 0 {
		t.Errorf("Expected 0, got %d", xor.GetOutputIdx("out"))
	}
	if xor.GetOutputIdx("outout") != -1 {
		t.Errorf("Expected -1, got %d", xor.GetOutputIdx("outout"))
	}
}

func TestCompiledChip_Name(t *testing.T) {
	if createXor().Name() != "Xor" {
		t.Errorf("Expected Xor, got %s", createXor().Name())
	}
}

func TestCompiledChip_SetInputEvaluateGetOutput(t *testing.T) {

	xor := createXor()
	xor.SetInput(1, true)
	xor.Evaluate()
	if xor.GetOutput(0) != true {
		t.Errorf("Expected 1, got: %d", 0)
	}
}
