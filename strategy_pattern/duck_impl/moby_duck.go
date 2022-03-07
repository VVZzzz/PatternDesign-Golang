package duck_impl

import "github.com/PatternDesign-Golang/strategy_pattern/duck_model"

func NewMobyDuck() *duck_model.Duck {
	mobyDuck := &duck_model.Duck{}
	mobyDuck.SetQuackBehavior(NewQuackWithLowValBehavior())
	mobyDuck.SetFlyBehavior(NewFlyWithZzBehavior())
	return mobyDuck
}
