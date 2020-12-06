// What is the highest seat ID on a boarding pass?
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

func OpennSlice(s string) []string {
  data, err := ioutil.ReadFile(s)
  check(err)
  sliced_data := strings.Split(string(data), "\n")
  if len(sliced_data[len(sliced_data) - 1]) == 0 {
    sliced_data = sliced_data[:len(sliced_data) - 1]
  }
  return sliced_data
}

func main() {
  highest := 0
  data := OpennSlice("./input")

  for _, s := range data {
    res := Parser(s)
    id := CalculateId(res[0], res[1])
    if id > highest {
      highest = id
    }
  }
  fmt.Printf("Highest seat ID on a boarding pass in the list: %d", highest)
}

func Parser(s string) []int {
  rowl := 0
  rowh := 128
  coll := 0
  colh := 8
  rowrule := s[:7]
  colrule := s[7:]
  rowdelta := 0
  coldelta := 0

  for _, c := range rowrule {
    fmt.Println(string(c))
    rowdelta = (rowh - rowl) / 2
    if string(c) == "F" {
      rowh = rowh - rowdelta
    } else {
      rowl = rowl + rowdelta
    }
    fmt.Printf("%d %d delta: %d\n", rowl, rowh, rowdelta)
  }

  for _, c := range colrule {
    coldelta = (colh - coll) / 2
    if string(c) == "L" {
      colh = colh - coldelta
    } else {
      coll = coll + coldelta
    }
    fmt.Printf("%d %d delta: %d\n", coll, colh, coldelta)
  }
  return []int{ rowl, coll }
}

func CalculateId(row int, col int) int {
  return row * 8 + col
}
