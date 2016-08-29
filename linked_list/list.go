package linked_list

type List interface {
	Search(value int) *Node
	Insert(value int)
	Delete()
	Size() int
	Head() *Node
	Tail() *Node
}

func NewAtomicList() *AtomicList {
	return &AtomicList{
		head:  nil,
		tail:  nil,
		count: 0,
	}
}

func NewBlockingList() *BlockingList {
	sentinel := &Node{
		Prev:  nil,
		Next:  nil,
		Value: -1,
	}

	return &BlockingList{
		head:  sentinel,
		tail:  sentinel,
		count: 0,
	}
}
