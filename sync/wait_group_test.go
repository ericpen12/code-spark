package sync

import (
	"sync"
	"testing"
	"time"
)

func TestUseWaitGroup(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		time.Sleep(5 * time.Second)
		wg.Done()

	}()
	go func() {
		wg.Wait()
		t.Log(1)
	}()
	wg.Wait()
	t.Log(2)
	time.Sleep(time.Second)
}
