package pattern

import (
	"fmt"
	"testing"
)

func Test_Pattern(t *testing.T) {
	context := NewContext(new(AddStrategy))
	fmt.Println(context.Operation(1, 2))
}