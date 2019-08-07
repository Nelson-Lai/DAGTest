package node

import "fmt"

type Node struct {
	parent   *Node
	children []*Node
	value    int
}

func (n *Node) Children(children ...*Node) {
	for _, node := range children {
		n.children = append(n.children, node)
	}
}

func (n *Node) PrintDepthFirst() {
	fmt.Println(n.value)
	if len(n.children) == 0 {
		return
	}
	for _, node := range n.children {
		node.PrintDepthFirst()
	}
}

func New(value int) *Node {
	n := new(Node)
	n.value = value
	return n
}
