package decorator_pattern

// MochaDecorator Mocha调料装饰者
// MochaDecorator 本质仍是一个 Coffe 所以仍可以被装饰,这也是 Decorator 的核心
type MochaDecorator struct {
	Coffee
}

func (m *MochaDecorator) Cost() float64 {
	return m.Coffee.Cost() + 0.02
}

func (m *MochaDecorator) GetDescription() string {
	return m.Coffee.GetDescription() + "Mocha"
}

type WhipDecorator struct {
	Coffee
}

func (m *WhipDecorator) Cost() float64 {
	return m.Coffee.Cost() + 0.03
}

func (m *WhipDecorator) GetDescription() string {
	return m.Coffee.GetDescription() + "Whip"
}

type MilkDecorator struct {
	Coffee
}

func (m *MilkDecorator) Cost() float64 {
	return m.Coffee.Cost() + 0.04
}

func (m *MilkDecorator) GetDescription() string {
	return m.Coffee.GetDescription() + "Milk"
}
