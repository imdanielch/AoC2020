package main

import (
  "testing"
)

// (Birth Year) - four digits; at least 1920 and at most 2002.
func TestValidateBYR(t *testing.T) {
  tables := []struct{
    x string
    y bool
  }{
    {"", false},
    {"0", false},
    {"2", false},
    {"1920", true},
    {"1978", true},
    {"2002", true},
    {"2020", false},
    {"2049", false},
    {"x321", false},
  }

  for _, table := range tables {
    res, err := ValidateBYR(table.x)
    if err != nil {
      t.Logf("error: \n\t%s\n", err)
      if res == true {
        t.Errorf("%s validated to %v, but should be %v", table.x, res, table.y)
      }
    }

    if res != table.y {
      t.Errorf("%s validated to %v, but should be %v", table.x, res, table.y)
    }
  }
}

// (Issue Year) - four digits; at least 2010 and at most 2020.
func TestValidateIYR(t *testing.T) {
  tables := []struct{
    x string
    y bool
  }{
    {"", false},
    {"2", false},
    {"2002", false},
    {"2010", true},
    {"2020", true},
    {"2049", false},
    {"x321", false},
  }

  for _, table := range tables {
    res, err := ValidateIYR(table.x)
    if err != nil {
      t.Logf("error: \n\t%s\n", err)
      if res == true {
        t.Errorf("%s validated to %v, but should be %v", table.x, res, table.y)
      }
    }

    if res != table.y {
      t.Errorf("%s validated to %v, but should be %v", table.x, res, table.y)
    }
  }
}

// (Expiration Year) - four digits; at least 2020 and at most 2030.
func TestValidateEYR(t *testing.T) {
  tables := []struct{
    x string
    y bool
  }{
    {"", false},
    {"2", false},
    {"2010", false},
    {"2020", true},
    {"2030", true},
    {"2049", false},
    {"x321", false},
  }

  for _, table := range tables {
    res, err := ValidateEYR(table.x)
    if err != nil {
      t.Logf("error: \n\t%s\n", err)
      if res == true {
        t.Errorf("%s validated to %v, but should be %v", table.x, res, table.y)
      }
    }

    if res != table.y {
      t.Errorf("%s validated to %v, but should be %v", table.x, res, table.y)
    }
  }
}

// (Height) - a number followed by either cm or in:
// - If cm, the number must be at least 150 and at most 193.
// - If in, the number must be at least 59 and at most 76.
func TestValidateHGT(t *testing.T) {
  tables := []struct {
    x string
    y bool
  }{
    {"60in", true},
    {"190cm", true},
    {"190in", false},
    {"190", false},
    {"", false},
    {"x3ds", false},
  }

  for _, table := range tables {
    res, err := ValidateHGT(table.x)
    if err != nil {
      t.Logf("error: \n\t%s\n", err)
      if res == true {
        t.Errorf("%s validated to %v, but should be %v", table.x, res, table.y)
      }
    }

    if res != table.y {
      t.Errorf("%s validated to %v, but should be %v", table.x, res, table.y)
    }
  }
}

// (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
func TestValidateHCL(t *testing.T) {
  tables := []struct {
    x string
    y bool
  }{
    {"#123abc", true},
    {"#123abz", false},
    {"", false},
    {"$123abc", false},
    {"123abc", false},
  }

  for _, table := range tables {
    res, err := ValidateHCL(table.x)
    if err != nil {
      t.Logf("error: \n\t%s\n", err)
      if res == true {
        t.Errorf("%s validated to %v, but should be %v", table.x, res, table.y)
      }
    }

    if res != table.y {
      t.Errorf("%s validated to %v, but should be %v", table.x, res, table.y)
    }
  }
}

// (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
func TestValidateECL(t *testing.T) {
  tables := []struct {
    x string
    y bool
  }{
    {"amb", true},
    {"blu", true},
    {"gry", true},
    {"grn", true},
    {"hzl", true},
    {"oth", true},
    {"brn", true},
    {"", false},
    {"@", false},
    {"wat", false},
  }

  for _, table := range tables {
    res, err := ValidateECL(table.x)
    if err != nil {
      t.Logf("error: \n\t%s\n", err)
      if res == true {
        t.Errorf("%s validated to %v, but should be %v", table.x, res, table.y)
      }
    }

    if res != table.y {
      t.Errorf("%s validated to %v, should be %v", table.x, res, table.y)
    }
  }
}

// (Passport ID) - a nine-digit number, including leading zeroes.

func TestValidatePID(t *testing.T) {
  tables := []struct {
    x string
    y bool
  }{
    {"000000001", true},
    {"123456789", true},
    {"12@456789", false},
    {"0123456789", false},
    {"01", false},
    {"", false},
    {"%", false},
  }

  for _, table := range tables {
    res, err := ValidatePID(table.x)
    if err != nil {
      t.Logf("error: \n\t%s\n", err)
      if res == true {
        t.Errorf("%s validated to %v, but should be %v", table.x, res, table.y)
      }
    }

    if res != table.y {
      t.Errorf("%s validated to %v, should be %v", table.x, res, table.y)
    }
  }
}
