package linked_list

import (
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

const N = 100

func testInsert(t *testing.T, list List) {
	var wg sync.WaitGroup
	wg.Add(N)

	for i := 0; i < N; i++ {
		go func() {
			list.Insert(i)
			wg.Done()
		}()
	}

	wg.Wait()
	assert.Equal(t, list.Size(), N)
	i := 0

	for cur := list.Head(); cur != nil; cur = cur.Next {
		assert.Equal(t, cur.Value, i)
		i++
	}
}

func testDelete(t *testing.T, list List) {
	for i := 0; i < N; i++ {
		list.Insert(i)
	}
	var wg sync.WaitGroup
	wg.Add(N)
	assert.Equal(t, list.Size(), N)

	for i := 0; i < N; i++ {
		go func() {
			list.Delete()
			wg.Done()
		}()
	}

	wg.Wait()
	assert.Equal(t, list.Size(), 0)
}

func testInsertDelete(t *testing.T, list List) {
	var wg sync.WaitGroup
	wg.Add(N * 2)

	for i := 0; i < N; i++ {
		go func() {
			list.Insert(i)
		}()
	}

	for i := 0; i < N; i++ {
		go func() {
			list.Delete()
		}()
	}

	wg.Wait()
	assert.Equal(t, list.Size(), 0)
}

func TestInsertAtomic(t *testing.T) {
	list := NewAtomicList()
	testInsert(t, list)
}

func TestInsertBlock(t *testing.T) {
	list := NewBlockingList()
	testInsert(t, list)
}

func TestDeleteAtomic(t *testing.T) {
	list := NewAtomicList()
	testDelete(t, list)
}

func TestDeleteBlock(t *testing.T) {
	list := NewBlockingList()
	testDelete(t, list)
}

func TestDeleteInsertAtomic(t *testing.T) {
	list := NewAtomicList()
	testInsertDelete(t, list)
}

func TestDeleteInsertBlock(t *testing.T) {
	list := NewBlockingList()
	testInsertDelete(t, list)
}
