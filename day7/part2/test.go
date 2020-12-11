package main

import "fmt"

func main() {
  data := make(map[string]map[string]int{})

  data["a"] = map[string]int{}
  data["a"]["w"] = 3

  fmt.Println(data)
}
