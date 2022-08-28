package decorator_pattern

import "fmt"

func TestDecorator() {
	fmt.Printf("Decorator Pattern Start...\n")
	// 普通咖啡
	generalCoffee := &GeneralCoffeeImpl{}
	// 普通咖啡+摩卡
	resultCoffe1 := &MochaDecorator{Coffee: generalCoffee}
	// 普通咖啡+摩卡+whip
	resultCoffe2 := &WhipDecorator{Coffee: resultCoffe1}
	fmt.Println(resultCoffe2.GetDescription())

	// 浓缩咖啡+Whip+Milk
	espressCoffee := &MilkDecorator{Coffee: &WhipDecorator{Coffee: &EspressoCoffeeImpl{}}}
	fmt.Println(espressCoffee.GetDescription())
	fmt.Printf("Decorator Pattern End...\n\n")
}
