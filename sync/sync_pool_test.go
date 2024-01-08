package sync

import (
	"sync"
	"testing"
)

type person struct {
	Name string
}

func TestSyncPool(t *testing.T) {
	pool := sync.Pool{
		New: func() any {
			return &person{}
		},
	}
	a := &person{
		Name: "tom",
	}
	pool.Put(a)
	pool.Put(a)
	_ = pool.Get().(*person)
	//t.Log(a2.Name)
}

func TestQueue(t *testing.T) {
	q := &queue{}
	for i := 0; i < 5; i++ {
		q.enQueue(1)
		q.info("入队")
	}
	for i := 0; i < 5; i++ {
		q.deQueue()
		q.info("出队")
	}
}
