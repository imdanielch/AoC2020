package main

import "testing"

func TestParseLoc(t *testing.T) {
  table := []struct {
    s string
    row int
    col int
    result int
  }{
    {"BFFFBBFRRR", 70, 7, 567},
    {"FFFBBBFRRR", 14, 7, 119},
    {"BBFFBBFRLL", 102, 4, 820},
    {"FBFBBFFRLR", 44, 5, 357},
  }

  for _, f := range table {
    res := Parser(f.s)
    if res[0] != f.row {
      t.Errorf("%s returned row %d, but %d is expected", f.s, res[0], f.row)
    }
    if res[1] != f.col {
      t.Errorf("%s returned column %d, but %d is expected", f.s, res[1], f.col)
    }
  }
}

func TestCalculateId(t *testing.T) {
  table := []struct {
    s string
    row int
    col int
    result int
  }{
    {"BFFFBBFRRR", 70, 7, 567},
    {"FFFBBBFRRR", 14, 7, 119},
    {"BBFFBBFRLL", 102, 4, 820},
    {"FBFBBFFRLR", 44, 5, 357},
  }
  for _, f := range table {
    res := CalculateId(f.row, f.col)
    if res != f.result {
      t.Errorf("%s returned %d, but %d is expected", f.s, res, f.result)
    }
  }
}
