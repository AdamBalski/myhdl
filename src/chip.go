package myhdl

type Chip interface {
	Name() string
	SetInput(idx int, value bool)
	GetInputIdx(inputName string) int
	GetOutputIdx(inputName string) int
	GetOutput(idx int) bool
	PropagateInputsToSubChips()
	Evaluate()
}
