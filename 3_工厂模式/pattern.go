package pattern

type Operation interface {
	Operation(float64, float64) float64
}

type AddOperation struct {
}

func (p *AddOperation) Operation(a, b float64) float64 {
	return a + b
}

type MultiOperation struct {
}

func (p *MultiOperation) Operation(a, b float64) float64 {
	return a * b
}

type Factory interface {
	GetOperation() Operation
}

type AddFactory struct {
}

func (p *AddFactory) GetOperation() Operation {
	return new(AddOperation)
}

type MultiFactory struct {
}

func (p *MultiFactory) GetOperation() Operation {
	return new(MultiOperation)
}
