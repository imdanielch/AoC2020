package main

import (
  "fmt"
  "log"
  "io/ioutil"
  "strconv"
  "strings"
  "regexp"
)

// helper function to check for errors
func check(e error) {
  if e != nil {
    log.Println(e)
  }
}

func main() {
  reqFields := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

  data, err := ioutil.ReadFile("./input")
  check(err)

  passports := strings.Split(string(data), "\n\n")

  count := 0
  for _, passport := range passports {
    invalid := false
    for i, _ := range reqFields {
      if strings.Contains(passport, reqFields[i]) == false {
        invalid = true
        break
      } else {
        val, err := ValidatePass(passport)
        check(err)
        if val == false {
          invalid = true
          break
        }
      }
    }
    if !invalid {
      count++
    }
  }

  fmt.Println("Count Total: ", count)
}

func ValidatePass (sl string) (bool, error) {
  valid := true
  cat := strings.Fields(sl)
  // create a map for passport
  mapy := make(map[string]string)
  for _, s := range cat {
    kv := strings.Split(s, ":")
    mapy[kv[0]] = kv[1]
  }
  table := []struct {
    sym string
    f func(string) (bool, error)
  }{
    {"byr", ValidateBYR},
    {"iyr", ValidateIYR},
    {"eyr", ValidateEYR},
    {"hgt", ValidateHGT},
    {"hcl", ValidateHCL},
    {"ecl", ValidateECL},
    {"pid", ValidatePID},
  }
  for _, field := range table {
    val, err := field.f(mapy[field.sym])
    check(err)
    if !val {
      valid = false
    }
  }
  return valid, nil
}

func ValidateBYR (s string) (bool, error) {
  i, err := strconv.Atoi(s)
  if err != nil {
    return false, err
  }

  if i >= 1920 && i <= 2002 {
    return true, nil
  }
  return false, nil
}

func ValidateIYR (s string) (bool, error) {
  i, err := strconv.Atoi(s)
  if err != nil {
    return false, err
  }
  if i >= 2010 && i <= 2020 {
    return true, nil
  }
  return false, nil
}

func ValidateEYR (s string) (bool, error) {
  i, err := strconv.Atoi(s)
  if err != nil {
    return false, err
  }
  if i >= 2020 && i <= 2030 {
    return true, nil
  }
  return false, nil
}

func ValidateHGT (s string) (bool, error) {
  if strings.Contains(s, "cm") {
    i, err := strconv.Atoi(s[:len(s) - 2])
    if err != nil {
      return false, err
    }

    if i >= 150 && i <= 193 {
      return true, nil
    }
  } else if strings.Contains(s, "in") {
    i, err := strconv.Atoi(s[:len(s) - 2])
    if err != nil {
      return false, err
    }

    if i >= 59 && i <= 76 {
      return true, nil
    }
  }
  return false, nil
}

func ValidateHCL (s string) (bool, error) {
  if len(s) > 0 {
    if s[0] == byte('#') {
      return regexp.Match("^[a-fA-F0-9]*$", []byte(s[1:]))
    }
  }
  return false, nil
}

func ValidateECL (s string) (bool, error) {
  validColors := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
  if Contains(validColors, s) {
    return true, nil
  }
  return false, nil
}

func ValidatePID(s string) (bool, error) {
  if len(s) == 9 {
    return regexp.Match("^[0-9]*$", []byte(s))
  }
  return false, nil
}

// checks if int slice contains the string in question
func Contains(slice []string, val string) (bool) {
  for _, entry := range slice {
    if entry == val {
      return true
    }
  }
  return false
}
