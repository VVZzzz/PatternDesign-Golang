package duck_impl

import "fmt"

type QuackWithLowValBehavior struct {
}

func (q *QuackWithLowValBehavior) Quack() {
	fmt.Println("Quack with Low Val")
}

func NewQuackWithLowValBehavior() *QuackWithLowValBehavior {
	return &QuackWithLowValBehavior{}
}
