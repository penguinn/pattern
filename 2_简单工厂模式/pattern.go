package pattern

import (
	"errors"
)

// 又称为静态工厂方法(Static Factory Method)模式，它属于类创建型模式。在简单工厂模式中，可以根据参数的不同返回不同类的实例。
// 简单工厂模式专门定义一个类来负责创建其他类的实例，被创建的实例通常都具有共同的父类。

type Operation interface {
	Operation(float64, float64) float64
}

type AddOperation struct {
}

func (p *AddOperation) Operation(x, y float64) float64 {
	return x + y
}

type MultiOperation struct {
}

func (p *MultiOperation) Operation(x, y float64) float64 {
	return x * y
}

type OperationFactory struct {
}

func (p *OperationFactory) CreateOperation(flag string) (Operation, error) {
	switch flag {
	case "+":
		return &AddOperation{}, nil
	case "*":
		return &MultiOperation{}, nil
	default:
		return nil, errors.New("不支持的运算符")
	}
}
