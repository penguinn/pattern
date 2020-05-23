package pattern

import (
	"testing"
)

func Test_Pattern(t *testing.T) {
	var person Person
	person.template = new(XiaoMing)
	person.Show()

	person.template = new(XiaoHong)
	person.Show()
}
