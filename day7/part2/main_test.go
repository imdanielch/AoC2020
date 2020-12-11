package main

import (
  "fmt"
  "testing"
)


func TestGetColors(t *testing.T) {
  table := []struct {
    s string
    v []string
  }{
    {
      " 5 shiny red bag, 2 muted blue bags.",
      []string{"shiny red", "muted blue"},
    },
    {
      " no other bags.",
      []string{},
    },
    {
      "",
      []string{},
    },

  }
  for _, x := range table {
    res := GetColors(x.s)
    if !SliceIsEq(res, x.v) {
      t.Errorf("\nGot: %v,\nWanted: %v", res, x.v)
    }

  }
}

func SliceIsEq( a, b []string) bool {
  if (a == nil) != (b == nil) {
    fmt.Println("One string slice was nil, the other was not")
    return false
  }

  if len(a) != len(b) {
    fmt.Println("both slices are not of equal length")
    return false
  }

  for i := range(a) {
    if a[i] != b[i] {
      fmt.Printf("A: %s\nB: %s\nDo not match\n", a[i], b[i])
      return false
    }
  }

  return true
}
