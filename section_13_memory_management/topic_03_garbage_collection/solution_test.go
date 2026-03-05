package garbagecollection

import "testing"

func TestObjectPool(t *testing.T) {
	pool := NewObjectPool(2, 1024)
	if pool == nil {
		t.Fatal("nil")
	}
	buf1 := pool.Get()
	if buf1 == nil {
		t.Fatal("Get returned nil")
	}
	if len(buf1) != 1024 {
		t.Errorf("buf len = %d", len(buf1))
	}
	pool.Put(buf1)
	buf2 := pool.Get()
	if buf2 == nil {
		t.Fatal("Get after Put returned nil")
	}
}

func TestProcessWithPool(t *testing.T) {
	pool := NewObjectPool(5, 256)
	reuses := ProcessWithPool(pool, 10)
	if reuses < 0 {
		t.Errorf("reuses = %d", reuses)
	}
}
