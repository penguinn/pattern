package pattern

import (
	"fmt"
)

// 定义一个操作中的算法的骨架，而将一些步骤延迟到到子类中。模板方法使得子类可以不改变一个算法的结构即可重定义该算法的某些特定步骤

type Template interface {
	Clothes()
	Pants()
}

type Person struct {
	template Template
}

func (p *Person) Clothes() {
	fmt.Println("clothes")
}

func (p *Person) Pants() {
	fmt.Println("pants")
}

func (p *Person) Show() {
	p.template.Clothes()
	p.template.Pants()
}

type XiaoMing struct {
	Person
}

func (p *XiaoMing) Clothes() {
	fmt.Println("t-shirt")
}

func (p *XiaoMing) Pants() {
	fmt.Println("jeans")
}

type XiaoHong struct {
	Person
}

func (p *XiaoHong) Clothes() {
	fmt.Println("tights")
}

func (p *XiaoHong) Pants() {
	fmt.Println("dress")
}
