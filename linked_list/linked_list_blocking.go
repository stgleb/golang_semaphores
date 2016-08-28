package linked_list

import "sync"

type BlockingList struct {
	head *Node
	tail *Node

	count int
	sync.RWMutex
}

func (list BlockingList) Search(value int) *Node {
	list.RLock()
	defer list.RUnlock()

	cur := list.head

	for cur != nil {
		if cur.Value == value {
			return cur
		}
	}

	return nil
}

func (list BlockingList) Insert(value int) {
	list.Lock()
	defer list.Unlock()

	newNode := &Node{
		Value: value,
		Next:  nil,
		Prev:  list.tail,
	}

	list.tail.Next = newNode
	list.count++
}

func (list BlockingList) Delete(value int) {
	list.Lock()
	defer list.Unlock()

	if list.tail == nil {
		return
	}

	list.tail = list.tail.Prev
	list.tail.Next = nil
}