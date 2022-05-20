package functions_test

import (
	"testing"

	"asdf/functions"
)

func TestMain(t *testing.T) {
	got := functions.MyFunc()
	want := true
	if got != want {
		t.Errorf("error: got %t want %t", got, want)
	}
}
