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

    if IsValid(l, h, sym, strings.TrimSpace(buf[1])) {
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
  valid := false
  if string(pass[l - 1]) == sym {
    valid = true
  }
  if string(pass[h - 1]) == sym {
    if !valid {
      valid = true
    } else {
      valid = false
    }

  }

  fmt.Printf("len(pass): %d", len(pass))
  fmt.Printf("\tsym:\t %s", sym)
  fmt.Printf("\tpass[%d]: %s", l-1, string(pass[l - 1]))
  fmt.Printf("\tpass[%d]: %s", h-1, string(pass[h - 1]))
  if valid {
    fmt.Printf("\tvalid: true")
  } else {
    fmt.Printf("\tvalid: false")
  }
  fmt.Printf("\n")
  return valid
}
