package mutex

import (
	"sync"
	"testing"
)

func TestSafeCounter(t *testing.T) {
	c := &SafeCounter{}
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() { defer wg.Done(); c.Increment() }()
	}
	wg.Wait()
	if got := c.Value(); got != 1000 {
		t.Errorf("Counter = %d; want 1000 (race condition?)", got)
	}
}

func TestSafeMap(t *testing.T) {
	m := NewSafeMap()
	if m == nil {
		t.Fatal("NewSafeMap returned nil")
	}
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			m.Set("key", i)
		}(i)
	}
	wg.Wait()
	_, ok := m.Get("key")
	if !ok {
		t.Error("key should exist")
	}
	if m.Len() != 1 {
		t.Errorf("Len = %d; want 1", m.Len())
	}
}

func TestSafeMap_NotFound(t *testing.T) {
	m := NewSafeMap()
	_, ok := m.Get("missing")
	if ok {
		t.Error("missing key should return false")
	}
}
