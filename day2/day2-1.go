// 1-3 a means that the password must contain a at least 1 time and at most 3 times.
// How many passwords are valid according to their policies?
package main

import (
  "fmt"
  "io/ioutil"
  "strconv"
  "strings"
)

// helper function to check for errors
func check(e error) {
  if e != nil {
    panic (e)
  }
}

func main() {
  data, err := ioutil.ReadFile("./input")
  check(err)

  // convert []byte to []string and split
  lines := strings.Split(string(data), "\n")

  valid_count := 0
  for _, line := range lines {
    // checking for EOF
    if len(line) == 0 {
      continue
    }

    buf := strings.Split(string(line), ":")
    //fmt.Println(i, s)

    l, h, sym := ParseRule(buf[0])
    if IsValid(l, h, sym, buf[1]) {
      valid_count++
    }
    //fmt.Printf("%d - %d: %s; %s\n", l, h, sym, strings.TrimSpace(buf[1]))
  }
  fmt.Printf("%d valid passwords", valid_count)

}

func ParseRule (buf string) (l int, h int, sym string) {
  spaced := strings.Fields(buf)
  count := strings.Split(spaced[0], "-")
  sym = spaced[1]

  l, err := strconv.Atoi(count[0])
  check(err)
  h, err2 := strconv.Atoi(count[1])
  check(err2)

  return l, h, sym
}

func IsValid (l int, h int, sym string, pass string) (bool) {
  counter := 0

  for _, char := range pass {
    if string(char) == sym {
      counter++
    }
  }

  if counter >= l && counter <= h {
    return true
  }

  return false
}
