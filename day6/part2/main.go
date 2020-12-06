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
  fmt.Println(Parse(data))
}

func Parse(data []string) int {
  results := 0
  for _, group := range data {
    list := make(map[string]int, 0)
    count := strings.Count(strings.TrimSpace(group), "\n")
    for _, a := range group{
      char := string(a)
      valid, _ := regexp.MatchString("[a-z]", char)
      if valid {
        list[char] = list[char] + 1
      }
    }
    fmt.Println("map: ", list)
    fmt.Println("count to match: ", count)
    for _, v := range list {
      if v == count + 1 {
        results++
      }
    }
  }
  return results
}
