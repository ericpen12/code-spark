package sync

import (
	"sync"
	"sync/atomic"
)

type Once struct {
	done uint32
	m    sync.Mutex
}

func (o *Once) Do(fn func()) {
	if atomic.LoadUint32(&o.done) != 0 {
		return
	}

	o.m.Lock()
	defer o.m.Unlock()
	if o.done == 0 {
		defer atomic.StoreUint32(&o.done, 1)
		fn()
	}
}
