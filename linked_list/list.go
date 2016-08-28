package linked_list

type List interface {
	Search(value int) *Node
	Insert(value int)
	Delete(value int)
}
