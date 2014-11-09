package main

type CompositeNode struct {
	Node
	Children []*Node
}

func NewCompositeNode(dirname string) CompositeNode {
	compositeNode := CompositeNode{Node: NewNode(dirname)}

	return compositeNode
}

func (composite_node *CompositeNode) AddChild(node *Node) error {
	composite_node.Children = append(composite_node.Children, node)

	return nil
}
