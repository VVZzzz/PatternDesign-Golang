package DuckModel

type Duck struct {
	// 自定义 Duck 外观等属性
	// 行为
	flyBehavior   FlyBehavior
	quackBehavior QuackBehavior
}

func (d *Duck) SetFlyBehavior(b FlyBehavior) {
	d.flyBehavior = b
}

func (d *Duck) SetQuackBehavior(b QuackBehavior) {
	d.quackBehavior = b
}

func (d *Duck) Fly() {
	d.flyBehavior.Fly()
}

func (d *Duck) Quack() {
	d.quackBehavior.Quack()
}
