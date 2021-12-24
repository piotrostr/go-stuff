package main 

import (
  "testing"
  // "github.com/stretchr/testify/assert"
)

func LongestConsecutiveString(strArr []string, k int) string {
  return strArr[0]
}

func TestLongestConsecutiveString(t *testing.T) {
  strArr := []string{"asdf"}
  out := LongestConsecutiveString(strArr, 1)
  if (out != "asdf") {
    t.Fail()
  }
}

// func main() {}
