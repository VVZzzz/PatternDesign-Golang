package main

import (
	"github.com/PatternDesign-Golang/decorator_pattern"
	"github.com/PatternDesign-Golang/factory_method"
	"github.com/PatternDesign-Golang/strategy_pattern/duck_test"
)

func main() {
	duck_test.Test()
	decorator_pattern.TestDecorator()
	factory_method.TestFactoryMethod()
}
