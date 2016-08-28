package linked_list

import "fmt"

type Node struct {
	Value int
	Next  *Node
	Prev  *Node
}

func (node Node) String() string {
	return fmt.Sprintf("Value %s", node.Value)
}
