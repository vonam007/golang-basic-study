package testfixtures

import (
	"os"
	"testing"
)

func TestTempFileHelper(t *testing.T) {
	path, err := TempFileHelper("hello world")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(path)
	content, err := ReadFileContent(path)
	if err != nil {
		t.Fatal(err)
	}
	if content != "hello world" {
		t.Errorf("content = %q", content)
	}
}

func TestSetupTestDir(t *testing.T) {
	dir, cleanup, err := SetupTestDir()
	if err != nil {
		t.Fatal(err)
	}
	if cleanup == nil {
		t.Fatal("cleanup is nil")
	}
	defer cleanup()
	info, err := os.Stat(dir)
	if err != nil {
		t.Fatal(err)
	}
	if !info.IsDir() {
		t.Error("should be directory")
	}
}
