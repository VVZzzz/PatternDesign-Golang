package adapter

// Target 需适配的目标接口
type Target interface {
	Request() string
}

// Adaptee 待适配的已有的接口
type Adaptee interface {
	SpecificRequest() string
}

// NewAdaptee 待适配接口的工厂函数
func NewAdaptee() Adaptee {
	return &adapteeImpl{}
}

// adapteeImpl 待适配接口实现
type adapteeImpl struct {
}

func (a *adapteeImpl) SpecificRequest() string {
	return "adapteeImpl SpecificRequest"
}

// Adapter 工厂方法
func NewAdapter(adaptee Adaptee) Target {
	return &adapter{adaptee}
}

// adapter 匿名组合 Adaptee 接口
type adapter struct {
	Adaptee
}

func (a *adapter) Request() string {
	return a.SpecificRequest()
}
