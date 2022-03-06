package DuckModel

type FlyBehavior interface {
	Fly() // interface 的方法需大写才可以被其他 package 实现
}

type QuackBehavior interface {
	Quack()
}
