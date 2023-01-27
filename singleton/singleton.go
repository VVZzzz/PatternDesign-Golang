package singleton

import "sync"

// Singleton 是单例模式接口，可导出的
// 通过该接口可以避免 GetInstance 返回一个包私有类型的指针
type Singleton interface {
	doSomething()
}

// 包私有类
type singleton struct{}

func (s *singleton) doSomething() {
}

var (
	instance *singleton
	once     sync.Once
)

// GetInstance 单例模式创建实例函数
func GetInstance() Singleton {
	once.Do(func() {
		instance = &singleton{}
	})
	return instance
}
