package pattern

import (
	"fmt"
)

// 动态地给一个对象添加一些额外的职责，就增加功能来说，装饰模式比生成子类更为灵活

type Component interface {
	Show()
}

type Person struct {
}

func (p *Person) Show() {
	fmt.Print("的企鹅\n")
}

type Clothes interface {
	Show()
	SetDecorator(component Component) Component
}

type TShirt struct {
	Component
}

func (p *TShirt) Show() {
	fmt.Print("T恤 ")
	p.Component.Show()
}

func (p *TShirt) SetDecorator(component Component) Component {
	p.Component = component
	return p
}

type Pants struct {
	Component
}

func (p *Pants) Show() {
	fmt.Print("牛仔裤 ")
	p.Component.Show()
}

func (p *Pants) SetDecorator(component Component) Component {
	p.Component = component
	return p
}
