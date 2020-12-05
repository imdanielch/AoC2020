package main

import (
  "fmt"
  "io/ioutil"
  "strings"
)

// helper function to check for errors
func check(e error) {
  if e != nil {
    panic (e)
  }
}

func main() {
  buf, err := ioutil.ReadFile("./input")
  check(err)
  count := 0
  iX := 0
  iY := 0
  // convert []byte to []string and split
  lines := strings.Split(string(buf), "\n")
  columns := len(lines[0])
  if len(lines[len(lines) - 1]) == 0 {
    lines = lines[:len(lines) - 1]
  }

  fmt.Printf("lines: %d, columns: %d\n", len(lines), columns)
  for iY < len(lines) {
    if iX >= columns {
      iX = iX % columns
    }

    fmt.Printf("%s", string(lines[iY][iX]))
    if lines[iY][iX] == byte('#') {
      count += 1
    }
    iX += 3
    iY += 1
  }
  fmt.Println("\nTotal count of trees = ", count)
}
