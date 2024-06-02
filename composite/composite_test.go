package composite

import "testing"

/*
+服装

	+男装
		-衬衣
		-夹克
	+女装
		-裙子
		-套装
*/
func TestComposite(t *testing.T) {
	// 组合对象
	root := NewComponent(CompositeNode, "服装")
	c1 := NewComponent(CompositeNode, "男装")
	c2 := NewComponent(CompositeNode, "女装")

	// 叶子对象
	leaf1 := NewComponent(LeafNode, "衬衣")
	leaf2 := NewComponent(LeafNode, "夹克")
	leaf3 := NewComponent(LeafNode, "裙子")
	leaf4 := NewComponent(LeafNode, "套装")

	// 树形结构组装
	root.AddChild(c1)
	root.AddChild(c2)
	c1.AddChild(leaf1)
	c1.AddChild(leaf2)
	c2.AddChild(leaf3)
	c2.AddChild(leaf4)

	root.PrintStruct()

}
