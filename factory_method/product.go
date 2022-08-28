package factory_method

import "fmt"

// Product
type PizzaInterface interface {
	Prepare()
	Bake()
	Cut()
	Box()
}

type Pizza struct {
	name string
}

func (p *Pizza) Prepare() {
	fmt.Printf("%s is preparing.\n", p.name)
}

func (p *Pizza) Bake() {
	fmt.Printf("%s is baking.\n", p.name)
}

func (p *Pizza) Cut() {
	fmt.Printf("%s is cutting.\n", p.name)
}

func (p *Pizza) Box() {
	fmt.Printf("%s is boxing.\n", p.name)
}

//NYStyle-Cheese Pizza
type NYCheesePizza struct {
	Pizza
}

func (p *NYCheesePizza) Cut() {
	fmt.Printf("%s is cutting into cheese style.", p.name)
}

func NewNYCheesePizza() *NYCheesePizza {
	return &NYCheesePizza{Pizza{name: "cheese"}}
}

//NYStyle-Milk Pizza
type NYMilkPizza struct {
	Pizza
}

func (p *NYMilkPizza) Cut() {
	fmt.Printf("%s is cutting into milk style.", p.name)
}

func NewNYMilkPizza() *NYMilkPizza {
	return &NYMilkPizza{Pizza{name: "milk"}}
}

//MiamiStyle-Milk Pizza
type MiamiMilkPizza struct {
	Pizza
}

func (p *MiamiMilkPizza) Box() {
	fmt.Printf("%s is boxing into miami style.", p.name)
}

func NewMiamiMilkPizza() *MiamiMilkPizza {
	return &MiamiMilkPizza{Pizza{name: "milk"}}
}
