package functions_test

import (
	"asdf/functions"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	got := functions.MyFunc()
	want := true
	assert.Equal(t, got, want)
}
