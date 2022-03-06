package DuckTest

import "github.com/PatternDesign-Golang/StrategyPattern/DuckImpl"

func Test() {
	mobyDuck := DuckImpl.NewMobyDuck()
	mobyDuck.Quack()
	mobyDuck.Fly()
}
