package main

import (
  "fmt"
  "io/ioutil"
  "strings"
  "regexp"
)

func OpennSlice(s string) []string {
  data, err := ioutil.ReadFile(s)
  if err != nil {
    panic (err)
  }
  sliced_data := strings.Split(string(data), "\n\n")
  if len(sliced_data[len(sliced_data) - 1]) == 0 {
    sliced_data = sliced_data[:len(sliced_data) - 1]
  }
  return sliced_data
}

func main() {
  data := OpennSlice("../input")
  results := 0
  for _, group := range data {
    list := make(map[string]bool, 0)
    for _, a := range group{
      char := string(a)
      valid, _ := regexp.MatchString("[a-z]", char)
      if valid && list[char] != true {
        list[char] = true
        results++
      }
    }
    fmt.Println(list)
  }
  fmt.Println(results)
}
