package cstrings

import (
	"strings"
	"testing"
)

func TestCount(t *testing.T) {
	count := strings.Count("cunkai11", "1")

	want := 2
	if count != want {
		t.Errorf(" count '%d' to want '%d'", count, want)
	}
}

func TestSplit(t *testing.T) {
	a := strings.Split("a,b,c", ",")
	b := []string{"a", "b", "c"}
	if len(a) != 3 {
		t.Errorf("a %q to b %q", a, b)
	}
}
