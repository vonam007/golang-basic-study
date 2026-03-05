package singleton

import (
	"sync"
	"testing"
)

func TestSingleton(t *testing.T) {
	ResetForTesting()
	db1 := GetDatabase()
	db2 := GetDatabase()
	if db1 == nil {
		t.Fatal("returned nil")
	}
	if db1 != db2 {
		t.Error("should return same instance")
	}
}

func TestSingleton_Concurrent(t *testing.T) {
	ResetForTesting()
	var wg sync.WaitGroup
	instances := make([]*Database, 100)
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()
			instances[idx] = GetDatabase()
		}(i)
	}
	wg.Wait()
	for i := 1; i < 100; i++ {
		if instances[i] != instances[0] {
			t.Fatal("not same instance")
		}
	}
}
