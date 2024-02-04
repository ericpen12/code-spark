package sync

import (
	"fmt"
	"net"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

// TestSyncOnce  sync.Once 的基本使用
func TestSyncOnce(t *testing.T) {
	var s sync.Once
	for i := 0; i < 10; i++ {
		go s.Do(func() {
			t.Log("Do")
		})
	}
	time.Sleep(time.Second)
}

type OnceA struct {
	done uint32
}

func (o *OnceA) Do(f func()) {
	if atomic.CompareAndSwapUint32(&o.done, 0, 1) {
		f()
	}
}

// TestOnceTheOtherImplement 测试源码中提到的有坑的实现方式
// 源码中提到的位置：/src/sync/once.go:51
// 问题现象：使用once 初始化资源，再使用再次使用资源时，该资源可能为 nil
// 问题原因：
// 1. 使用 atomic.CompareAndSwapUint32 作为只执行一次的判断条件时，done 会在 f 执行之前就被置为 1
// 2. 由于 done 已经被置为 1，所以再次调用 Do 时，会直接返回，不会执行 f
// 3. 如果在另外一个方法中，使用 Once 初始化资源， 这里不会执行 f，异步初始化资源的可能还没有结束，导致下文使用的资源是 nil
// ps：正确操作：如果异步初始化还没有完成，这里需要等待
func TestOnceTheOtherImplement(t *testing.T) {
	var once OnceA
	var conn net.Conn
	// 假设其他协程异步初始化资源，耗时比较长
	go func() {
		fun1 := func() {
			time.Sleep(5 * time.Second) //模拟初始化的速度很慢
			conn, _ = net.DialTimeout("tcp", "baidu.com:80", time.Second)
		}
		once.Do(fun1)
	}()
	// 当前协程初始化资源
	once.Do(func() {
		// 由于 done 在资源初始化之前就应该为1，所以这里不会执行 fun2
		t.Log("执行fun2")
		conn, _ = net.DialTimeout("tcp", "baidu.com:80", time.Second)
	})
	_, err := conn.Write([]byte("\"GET / HTTP/1.1\\r\\nHost: baidu.com\\r\\n Accept: */*\\r\\n\\r\\n\""))
	if err != nil {
		t.Log("err:", err)
	}
}

func TestOnce(t *testing.T) {
	s := &Once{}
	for i := 0; i < 100000; i++ {
		go func() {
			s.Do(func() {
				t.Log("once")
			})
		}()
	}
	time.Sleep(100 * time.Millisecond)
}
