package pattern

import (
	"fmt"
	"testing"
)

func Test_Pattern(t *testing.T) {
	operation := new(AddFactory).GetOperation()

	fmt.Println(operation.Operation(2, 4))
}
