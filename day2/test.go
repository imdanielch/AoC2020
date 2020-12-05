package main

import (
  "fmt"
  "strings"
)

func main() {
  data := []string{ "test1 1 1a", "test2 2 2b ", "test3  3   3c    " }
  for _, s := range data {
    fmt.Println("s: ", s)
    line := strings.Fields(s)
    for _, str := range line {
      fmt.Println("str: ", str)
    }
  }
}
