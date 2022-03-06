package DuckImpl

import "fmt"

type FlyWithZzBehavior struct {
}

func (f *FlyWithZzBehavior) Fly() {
	fmt.Println("Fly with ZZ")
}

func NewFlyWithZzBehavior() *FlyWithZzBehavior {
	return &FlyWithZzBehavior{}
}
