package composite

import "fmt"

// 抽象组件对象
type Component interface {
	Parent() Component
	SetParent(Component)
	Name() string
	SetName(string)
	AddChild(Component)
	PrintStruct()
}

// 抽象组件实体
type component struct {
	parent Component
	name   string
}

func (c *component) Parent() Component {
	return c.parent
}

func (c *component) SetParent(parent Component) {
	c.parent = parent
}

func (c *component) Name() string {
	return c.name
}

func (c *component) SetName(s string) {
	c.name = s
}

func (c *component) AddChild(child Component) {
}

func (c *component) PrintStruct() {
}

// 复合结构 or 叶子节点
const (
	LeafNode = iota
	CompositeNode
)

// 叶子节点
type Leaf struct {
	component
}

func (l *Leaf) PrintStruct() {
	fmt.Printf("leaf %s\n", l.name)
}

func NewLeaf() Component {
	return &Leaf{}
}

// 复合结构
type Composite struct {
	component
	childs []Component // 叶子
}

func (c *Composite) PrintStruct() {
	fmt.Printf("composite %s\n", c.name)
	for _, child := range c.childs {
		child.PrintStruct()
	}
}

func (c *Composite) AddChild(child Component) {
	child.SetParent(c)
	c.childs = append(c.childs, child)
}

func NewComposite() Component {
	return &Composite{
		childs: make([]Component, 0),
	}
}

func NewComponent(nodeType int, name string) Component {
	var c Component
	switch nodeType {
	case LeafNode:
		c = NewLeaf()
	case CompositeNode:
		c = NewComposite()
	}
	c.SetName(name)
	return c
}
