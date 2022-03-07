package duck_test

import "github.com/PatternDesign-Golang/strategy_pattern/duck_impl"

func Test() {
	mobyDuck := duck_impl.NewMobyDuck()
	mobyDuck.Quack()
	mobyDuck.Fly()
}
