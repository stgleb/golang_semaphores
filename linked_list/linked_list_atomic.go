package linked_list

import (
	"sync/atomic"
	"unsafe"
)

type AtomicList struct {
	head *Node
	tail *Node

	count int64
}

// Search across list
func (list AtomicList) Search(value int) *Node {
	cur := list.head
	var pos int64

	for cur != nil {
		if cur.Value == value {
			// If element was deleted after been found
			if atomic.AddInt64(&pos, -list.count) > 0 {
				return cur
			} else {
				return nil
			}
		} else {
			cur = cur.Next
		}
		pos++
	}

	return nil
}

// Insert value to the end of the list
func (list AtomicList) Insert(value int) {
	currentTail := list.tail

	node := &Node{
		Value: value,
		Prev:  list.tail,
		Next:  nil,
	}

	// Loop until tail changing successful
	for {
		if atomic.CompareAndSwapPointer(
			(*unsafe.Pointer)(unsafe.Pointer(list.tail)),
			unsafe.Pointer(currentTail),
			unsafe.Pointer(node)) {
			break
		}
	}

	atomic.AddInt64(&list.count, 1)
}

// Delete value from the end of the list
func (list AtomicList) Delete() {
	currentTail := list.tail
	newHead := list.tail.Prev

	// Loop until tail changing successful
	for {
		if atomic.CompareAndSwapPointer(
			(*unsafe.Pointer)(unsafe.Pointer(list.tail)),
			unsafe.Pointer(currentTail),
			unsafe.Pointer(newHead)) {
			break
		}
	}

	atomic.AddInt64(&list.count, -1)
}

func (list AtomicList) Size() int {
	return int(list.count)
}

func (list AtomicList) Head() *Node {
	return list.head
}

func (list AtomicList) Tail() *Node {
	return list.tail
}
