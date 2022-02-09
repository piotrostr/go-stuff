package main 

import (
  "testing"
)

func LongestString(strArr []string, k int) string {
  var longestSoFar string = strArr[0]
  for i := 1; i < len(strArr); i++ {
    if len(strArr[i]) > len(longestSoFar) {
      longestSoFar = strArr[i]
    }
  }
  return longestSoFar
}

func TestLongestString(t *testing.T) {
  strArr := []string{"asdf", "a", "as", "asd"}
  out := LongestString(strArr, 1)
  if out != "asdf" {
    t.Fail()
  }
}

func RemoveSpaces(str string) string {
  var newStr string = ""
  for i := 0; i < len(str); i++ {
    if string(str[i]) != string(" ") {
      newStr += string(str[i])
    }
  }
  return newStr
}

func TestRemoveSpaces(t *testing.T) {
  str := "8 j 8   mBliB8g  imjB8B8  jl  B"
  noSpacesStr := RemoveSpaces(str)
  if noSpacesStr != "8j8mBliB8gimjB8B8jlB" {
    t.Fail()
  }
}

func main() {}
