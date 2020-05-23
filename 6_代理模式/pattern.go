package pattern

import (
	"fmt"
)

// 为其他对象提供一种代理以控制对这个对象的访问

type Subject interface {
	Do()
}

type RealSubject struct {
}

func (p *RealSubject) Do() {
	fmt.Print("real ")
}

type Proxy struct {
	subject RealSubject
}

func (p *Proxy) Do() {
	fmt.Print("pre ")

	p.subject.Do()

	fmt.Print("after\n")
}
