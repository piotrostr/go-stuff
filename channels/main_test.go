package main

import (
	"testing"
)

func TestNothing(t *testing.T) {
	got := 0
	want := 1
	if got != want {
		t.Errorf("got %d; want %d", got, want)
	}
}
