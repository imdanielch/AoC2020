// Your seat wasn't at the very front or back, though; the seats with IDs +1 and -1 from yours will be in your list.
// What is the ID of your seat?
package main

import (
  "fmt"
  "io/ioutil"
  "strings"
  "sort"
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
  data := OpennSlice("./input")
  mySeat := 0
  seats := make([]int, 0, 0)

  for _, s := range data {
    res := Parser(s)
    id := CalculateId(res[0], res[1])
    seats = append(seats, id)
  }
  sort.Ints(seats)
  for i := 1; i < (len(seats) - 1); i++ {
     if seats[i - 1] - seats[i] != -1 {
      mySeat = seats[i] - 1
      if seats[i - 1] + 1 == mySeat {
        fmt.Printf("My seat ID: %d; ID before me: %d; ID after me %d\n", mySeat, seats[i - 1], seats[i])
      }
    }
  }
  fmt.Printf("My seat ID: %d", mySeat)
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
    rowdelta = (rowh - rowl) / 2
    if string(c) == "F" {
      rowh = rowh - rowdelta
    } else {
      rowl = rowl + rowdelta
    }
  }

  for _, c := range colrule {
    coldelta = (colh - coll) / 2
    if string(c) == "L" {
      colh = colh - coldelta
    } else {
      coll = coll + coldelta
    }
  }
  return []int{ rowl, coll }
}

func CalculateId(row int, col int) int {
  return row * 8 + col
}
