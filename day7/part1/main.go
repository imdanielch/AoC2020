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
  sliced_data := strings.Split(string(data), "\n")
  if len(sliced_data[len(sliced_data) - 1]) == 0 {
    sliced_data = sliced_data[:len(sliced_data) - 1]
  }
  return sliced_data
}

func main() {
  data := OpennSlice("../input")
  init := []string{"shiny gold"}
  res := []string{}
  rules := make(map[string][]string)
  // create map where key = color bag, value = colors which can contain the key
  for _, line := range data {
    slice := strings.Split(line, " bags contain ")
    colors := GetColors(slice[1])
    for _, c := range colors {
      rules[c] = append(rules[c], slice[0])
    }
  }
  results := Parse(rules, res, init)

  fmt.Println(results)
  fmt.Println(len(results))
}

func Parse(rules map[string][]string, r []string, n []string) []string {
  if len(n) == 0 {
    return r
  }
  next := []string{}
  for _, e := range n {
    if len(rules[e]) > 0 {
      for _, c := range rules[e] {
        if !Contains(r, c) {
          next = append(next, c)
          r = append(r, c)
        }
      }
    }
  }
  return Parse(rules, r, next)
}

func GetColors(s string) []string {
  colors := make([]string, 0)
  re := regexp.MustCompile(`\d (?P<color>[a-z\s]*) bag`)
  vals := re.FindAllStringSubmatch(s, -1)
  for _, a := range vals {
    i := re.SubexpIndex("color")
    colors = append(colors, a[i])
  }
  return colors
}

func Contains(slice []string, val string) (bool) {
  for _, entry := range slice {
    if entry == val {
      return true
    }
  }
  return false
}
