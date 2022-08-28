package factory_method

import "fmt"

func TestFactoryMethod() {
	fmt.Printf("Factory Method Pattern Start...\n")
	nyPizzaStore := NewNYPizzaStore()
	nyPizzaStore.CreatePizza("cheese")
	nyCheesePizza := nyPizzaStore.OrderPizza()
	fmt.Println("ordered a NYStyle Cheese Pizza ", nyCheesePizza)

	miamiPizzaStore := NewMiamiPizzaStore()
	miamiPizzaStore.CreatePizza("milk")
	miamiMilkPizza := miamiPizzaStore.OrderPizza()
	fmt.Println("ordered a MiamiStyle Milk Pizza ", miamiMilkPizza)
	fmt.Printf("Factory Method Pattern End...\n\n")
}
