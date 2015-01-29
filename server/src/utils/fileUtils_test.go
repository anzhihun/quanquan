package utils

import (
	"path/filepath"
	"testing"
)

func TestMakePath(t *testing.T) {
	path := MakePath(nil)
	expected := ""
	if path != expected {
		t.Fatal("failed to make path. expect: ", expected, " actual: ", path)
	}

	path = MakePath([]string{})
	expected = ""
	if path != expected {
		t.Fatal("failed to make path. expect: ", expected, " actual: ", path)
	}

	path = MakePath([]string{"1"})
	expected = "1"
	if path != expected {
		t.Fatal("failed to make path. expect: ", expected, " actual: ", path)
	}

	path = MakePath([]string{"1", "2"})
	expected = "1" + string(filepath.Separator) + "2"
	if path != expected {
		t.Fatal("failed to make path. expect: ", expected, " actual: ", path)
	}

	path = MakePath([]string{"1", "2", "3"})
	expected = "1" + string(filepath.Separator) + "2" + string(filepath.Separator) + "3"
	if path != expected {
		t.Fatal("failed to make path. expect: ", expected, " actual: ", path)
	}
}
