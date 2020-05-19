package pattern

// 它定义了算法家族，分别封装起来，让他们之间可以互相替换，此模式让算法的变化，不会影响到使用算法的客户

type Strategy interface {
	Operation(float64, float64) float64
}

type AddStrategy struct {
}

func (p *AddStrategy) Operation(x, y float64) float64 {
	return x + y
}

type MultiStrategy struct {
}

func (p *MultiStrategy) Operation(x, y float64) float64 {
	return x * y
}

type Context struct {
	Strategy
}

func NewContext(strategy Strategy) *Context {
	return &Context{strategy}
}

func (p *Context) Operation(x, y float64) float64 {
	return p.Strategy.Operation(x, y)
}
