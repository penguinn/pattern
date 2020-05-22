package pattern

import (
	"fmt"
)

type OperationController interface {
	Control()
}

type AndroidOperationController struct {
}

func (p *AndroidOperationController) Control() {
	fmt.Println("android operation controller")
}

type IosOperationController struct {
}

func (p *IosOperationController) Control() {
	fmt.Println("ios operation controller")
}

type UIController interface {
	Display()
}

type AndroidUIController struct {
}

func (p *AndroidUIController) Display() {
	fmt.Println("android ui controller")
}

type IosUIController struct {
}

func (p *IosUIController) Display() {
	fmt.Println("ios ui controller")
}

type Factory interface {
	GenerateOperation() OperationController
	GenerateUI() UIController
}

type IosFactory struct {
}

func (p *IosFactory) GenerateOperation() OperationController {
	return new(IosOperationController)
}

func (p *IosFactory) GenerateUI() UIController {
	return new(IosUIController)
}

type AndroidFactory struct {
}

func (p *AndroidFactory) GenerateOperation() OperationController {
	return new(AndroidOperationController)
}

func (p *AndroidFactory) GenerateUI() UIController {
	return new(AndroidUIController)
}
