package interfacecomposition

import "testing"

func TestNewMemoryStore(t *testing.T) {
	s := NewMemoryStore()
	if s == nil {
		t.Fatal("NewMemoryStore returned nil")
	}
}

func TestSaveAndLoad(t *testing.T) {
	s := NewMemoryStore()
	if err := s.Save("hello"); err != nil {
		t.Fatalf("Save error: %v", err)
	}
	got, err := s.Load("hello")
	if err != nil {
		t.Fatalf("Load error: %v", err)
	}
	if got != "hello" {
		t.Errorf("Load = %q; want hello", got)
	}
}

func TestLoadNotFound(t *testing.T) {
	s := NewMemoryStore()
	_, err := s.Load("missing")
	if err == nil {
		t.Error("Load missing key should return error")
	}
}

func TestDelete(t *testing.T) {
	s := NewMemoryStore()
	s.Save("test")
	if err := s.Delete("test"); err != nil {
		t.Fatalf("Delete error: %v", err)
	}
	_, err := s.Load("test")
	if err == nil {
		t.Error("After delete, Load should fail")
	}
	if err := s.Delete("missing"); err == nil {
		t.Error("Delete missing key should return error")
	}
}

func TestInterfaceSatisfaction(t *testing.T) {
	s := NewMemoryStore()
	var _ Storage = s     // MemoryStore satisfies Storage
	var _ FullStorage = s // MemoryStore satisfies FullStorage
	var _ Saver = s       // MemoryStore satisfies Saver
	var _ Loader = s      // MemoryStore satisfies Loader
	var _ Deleter = s     // MemoryStore satisfies Deleter
}

func TestCopyData(t *testing.T) {
	src := NewMemoryStore()
	dst := NewMemoryStore()
	src.Save("key1")
	src.Save("key2")
	err := CopyData(src, dst, []string{"key1", "key2"})
	if err != nil {
		t.Fatalf("CopyData error: %v", err)
	}
	v, err := dst.Load("key1")
	if err != nil || v != "key1" {
		t.Errorf("dst.Load(key1) = (%q, %v); want (key1, nil)", v, err)
	}
}
