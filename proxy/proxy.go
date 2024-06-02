package proxy

type Subject interface {
	Do() string
}

// RealSubject 真实对象
type RealSubject struct {
}

func (r *RealSubject) Do() string {
	return "real"
}

// Proxy RealSubject 真实对象的代理层
type Proxy struct {
	realSubject RealSubject
}

func (receiver *Proxy) Do() string {
	var res string

	// 在调用真实对象之前的工作，检查缓存，判断权限，实例化真实对象等。
	res += "proxy start "

	// 调用真实对象
	res += receiver.realSubject.Do()

	// 调用之后的操作，如缓存结果，对结果进行处理等。
	res += " proxy end"
	return res
}
