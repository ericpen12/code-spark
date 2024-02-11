package sync

import (
	"sync"
	"testing"
)

func TestStore(t *testing.T) {
	var m sync.Map
	m.Store(1, 1)
}

func TestLoad(t *testing.T) {
	var m sync.Map
	m.Store(1, 1)
	ret, ok := m.Load(1)
	t.Log(ret, ok)
}

func TestDelete(t *testing.T) {
	var m sync.Map
	m.Store(1, 1)
	m.Delete(1)
}
