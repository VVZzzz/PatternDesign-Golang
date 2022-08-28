package duck_test

import (
	"fmt"

	"github.com/PatternDesign-Golang/strategy_pattern/duck_impl"
)

func Test() {
	fmt.Printf("Strategy Pattern Start...\n")
	mobyDuck := duck_impl.NewMobyDuck()
	mobyDuck.Quack()
	mobyDuck.Fly()
	fmt.Printf("Strategy Pattern End...\n\n")
}
