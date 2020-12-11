package main

import (
  "log"
  "io/ioutil"
  "strings"
  "strconv"
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
  init := map[string]int{"shiny gold": 1}
  var rules = map[string]map[string]int{}
  for _, line := range data {
    slice := strings.Split(line, " bags contain ")
    inBags := GetNColors(slice[1])
    rules[slice[0]] = map[string]int{}
    for k, v := range inBags {
      log.Printf("\tinBags: k: %s; v: %d\n", k, v)
      rules[slice[0]][k] = v
      log.Printf("\trules[slice[0]]: %v\n", rules[slice[0]])
    }
  }
  log.Println("rules:\n", rules)
  results := Parse(rules, init)

  log.Println(results)
}

func Parse(rules map[string]map[string]int, n map[string]int) int {
  log.Println("n: ", n)
  result := 0
  // n = {"color": num, "color2": num... }
  for k, v := range n {
    log.Printf("\tk: %s, v: %d\n", k, v)
    log.Printf("\tlen(rules[k]): %d\n", len(rules[k]))
    if len(rules[k]) > 0 {
      log.Printf("\trules[k]: %v\n", rules[k])
      for s, d := range rules[k] { // for the color and digit of the upper bag
        next := map[string]int{}
        log.Printf("\t\ts: %s, d: %d\n", s, d)
        log.Printf("\t\tnext[s] = d: s: %s, d: %d\n", s, d)
        next[s] = d
        result = result + d + d * Parse(rules, next)
      }
    } else {
      log.Printf("\t--- end loop on %s\n---", k)
    }
  }
  return result
}

func CalcBags(n int, i int) int {
  return n + n * i
}

func GetNColors(s string) map[string]int {
  log.Printf("GetNColors: s: %s\n", s)
  entry := make(map[string]int)
  re := regexp.MustCompile(`(?P<num>\d) (?P<color>[a-z\s]*) bag`)
  vals := re.FindAllStringSubmatch(s, -1)
  log.Printf("GetNColors: vals: %v\n", vals)
  for _, a := range vals {
    log.Printf("GetNColors: a: %v\n", a)
    i := re.SubexpIndex("color")
    n := re.SubexpIndex("num")
    nint, err := strconv.Atoi(a[n])
    if err != nil {
      log.Println("error: ", err)
    }
    log.Printf("GetNColors: \ta[i]: %s\n", a[i])
    log.Printf("GetNColors: \tnint: %d\n", nint)
    entry[a[i]] = nint
  }
  log.Println("GetNColors: entry: ", entry)
  return entry
}
