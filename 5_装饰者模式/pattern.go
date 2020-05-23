package pattern

import (
	"fmt"
)

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
