package pattern

import (
	"testing"
)

func Test_Pattern(t *testing.T) {
	var person Component
	person = new(Person)
	person = new(TShirt).SetDecorator(person)
	person = new(Pants).SetDecorator(person)
	person.Show()
}
