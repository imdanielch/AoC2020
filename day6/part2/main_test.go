package main

import (
  "testing"
)

func TestParse(t *testing.T) {
  data := []string{ "abc", "a\nb\nc", "ab\nac", "a\na\na\na", "b"}
  result := Parse(data)
  if result != 6 {
    t.Errorf("got %d, expected 6", result)
  }
}
