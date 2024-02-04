package sync

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"
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

// TestSyncPoolRelease 观测 pool 释放空间
func TestSyncPoolRelease(t *testing.T) {
	pool := sync.Pool{
		New: func() any {
			return &person{}
		},
	}
	pool.Put(&person{Name: "tom"})
	r := reflect.ValueOf(pool)
	local := r.FieldByName("local")
	victim := r.FieldByName("victim")

	t.Logf("pool 刚存储，local=%v, victim=%v", local, victim)
	runtime.GC()
	r = reflect.ValueOf(pool)

	local = r.FieldByName("local")
	victim = r.FieldByName("victim")
	t.Logf("第一次GC, victim=local, local 置空，local=%v, victim=%v", local, victim)

	runtime.GC()
	r = reflect.ValueOf(pool)

	local = r.FieldByName("local")
	victim = r.FieldByName("victim")

	t.Logf("第二次GC, victim=local, local 置空，local=%v, victim=%v", local, victim)
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

func TestCode01(t *testing.T) {
	// 源码片段
	// go1.19/src/sync/pool.go:161
	//
	//for i := 0; i < int(size); i++ {
	//	l := indexLocal(locals, (pid+i+1)%int(size))
	//	if x, _ := l.shared.popTail(); x != nil {
	//		return x
	//	}
	//}

	size := 10
	pids := []int{
		// pid < size
		// 跟直接遍历不同的是，下标会遍历一圈，当前的 pid 在最后
		1,
		2,
		// pid > size 时，不会超出
		11,
		22,
	}
	for _, pid := range pids {
		var s []string
		for i := 0; i < size; i++ {
			index := (pid + i + 1) % size
			s = append(s, fmt.Sprintf("%d", index))
		}
		t.Logf("pid=%d, index=%s", pid, strings.Join(s, ","))
	}
}
