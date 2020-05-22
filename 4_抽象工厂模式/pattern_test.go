package pattern

import (
	"testing"
)

func Test_Pattern(t *testing.T) {
	factory := new(IosFactory)
	operation := factory.GenerateOperation()
	operation.Control()
	ui := factory.GenerateUI()
	ui.Display()
}