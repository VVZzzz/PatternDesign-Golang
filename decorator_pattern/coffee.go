package decorator_pattern

//Coffee 咖啡 interface(抽象类),被装饰者
type Coffee interface {
	Cost() float64
	GetDescription() string
}

// GeneralCoffeeImpl 普通咖啡组件
type GeneralCoffeeImpl struct {
}

func (b *GeneralCoffeeImpl) Cost() float64 {
	return 0.89
}

func (b *GeneralCoffeeImpl) GetDescription() string {
	return "普通咖啡"
}

// DarkRoastCoffeeImpl 深焙咖啡组件
type DarkRoastCoffeeImpl struct {
}

func (b *DarkRoastCoffeeImpl) Cost() float64 {
	return 0.99
}

func (b *DarkRoastCoffeeImpl) GetDescription() string {
	return "深焙咖啡"
}

// EspressoCoffeeImpl 深焙咖啡组件
type EspressoCoffeeImpl struct {
}

func (b *EspressoCoffeeImpl) Cost() float64 {
	return 0.99
}

func (b *EspressoCoffeeImpl) GetDescription() string {
	return "浓缩咖啡"
}
