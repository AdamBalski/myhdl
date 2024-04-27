package myhdl

type CompiledChip struct {
	name         string
	inputsNames  []string
	outputsNames []string
	inputs       *BitField
	outputs      *BitField
	table        []*BitField
}

func (chip *CompiledChip) Name() string {
	return chip.name
}

func (chip *CompiledChip) SetInput(idx int, value bool) {
	chip.inputs.Set(idx, value)
}

func (chip *CompiledChip) GetInputIdx(inputName string) int {
	for i, name := range chip.inputsNames {
		if name == inputName {
			return i
		}
	}
	return -1
}

func (chip *CompiledChip) GetOutputIdx(outputName string) int {
	for i, name := range chip.outputsNames {
		if name == outputName {
			return i
		}
	}
	return -1
}

func (chip *CompiledChip) GetOutput(idx int) bool {
	return chip.outputs.Get(idx)
}

func (chip *CompiledChip) PropagateInputsToSubChips() {
	// noop, because CompiledChips don't have any sub-chips
}

func (chip *CompiledChip) Evaluate() {
	chip.outputs = chip.table[chip.inputs.ToShort()]
}
