package node

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

func New(value int) *Node {
	n := new(Node)
	n.value = value
	return n
}
