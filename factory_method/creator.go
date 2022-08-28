package factory_method

import "fmt"

//Creator
type PizzaStoreInterface interface {
	// factoryMethod()
	CreatePizza(typeName string)
	// anOperation()
	OrderPizza() PizzaInterface
}

type PizzaStore struct {
	//PizzaStoreInterface
	pizza PizzaInterface
}

// OrderPizza anOperation()
func (p *PizzaStore) OrderPizza() PizzaInterface {
	if p.pizza == nil {
		fmt.Printf("Warning: wrong typeName %s\n")
		return nil
	}
	p.pizza.Prepare()
	p.pizza.Bake()
	p.pizza.Cut()
	p.pizza.Box()
	return p.pizza
}
