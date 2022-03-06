package DuckImpl

import "github.com/PatternDesign-Golang/StrategyPattern/DuckModel"

func NewMobyDuck() *DuckModel.Duck {
	mobyDuck := &DuckModel.Duck{}
	mobyDuck.SetQuackBehavior(NewQuackWithLowValBehavior())
	mobyDuck.SetFlyBehavior(NewFlyWithZzBehavior())
	return mobyDuck
}
