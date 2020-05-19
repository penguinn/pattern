package pattern

import (
	"fmt"
	"log"
	"testing"
)

func Test_Pattern(t *testing.T) {
	operation, err := new(OperationFactory).CreateOperation("+")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(operation.Operation(2, 4))
}
