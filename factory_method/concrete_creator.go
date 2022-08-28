package factory_method

//NewYork PizzaFactory
type NYPizzaStore struct {
	PizzaStore
}

func (p *NYPizzaStore) CreatePizza(typeName string) {
	switch typeName {
	case "cheese":
		p.PizzaStore.pizza = NewNYCheesePizza()
		return
	case "milk":
		p.PizzaStore.pizza = NewNYMilkPizza()
		return
	default:
		return
	}
}

// NYPizzaFactory 同样可覆盖默认operation函数

func NewNYPizzaStore() PizzaStoreInterface {
	return &NYPizzaStore{PizzaStore: PizzaStore{}}
}

// Miami PizzaFactory
type MiamiPizzaStore struct {
	PizzaStore
}

func (p *MiamiPizzaStore) CreatePizza(typeName string) {
	switch typeName {
	case "milk":
		p.PizzaStore.pizza = NewMiamiMilkPizza()
		return
	default:
		return
	}
}

func NewMiamiPizzaStore() PizzaStoreInterface {
	return &MiamiPizzaStore{PizzaStore: PizzaStore{}}
}
