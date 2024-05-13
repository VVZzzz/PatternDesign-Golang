package prototype

// Cloneable 原型对象需要实现的接口
type Cloneable interface {
	Clone() Cloneable
}

// PrototypeManager 原型管理器
type PrototypeManager struct {
	prototypes map[string]Cloneable
}

func NewPrototypeManager() *PrototypeManager {
	return &PrototypeManager{prototypes: make(map[string]Cloneable)}
}

func (p *PrototypeManager) Set(key string, prototype Cloneable) {
	p.prototypes[key] = prototype
}

func (p *PrototypeManager) Get(key string) Cloneable {
	return p.prototypes[key].Clone()
}
