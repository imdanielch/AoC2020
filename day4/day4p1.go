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
  reqFields := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

  data, err := ioutil.ReadFile("./input")
  check(err)

  passports := strings.Split(string(data), "\n\n")

  count := 0
  for _, passport := range passports {
    // create a map for passport
    invalid := false
    for i, _ := range reqFields {
      if strings.Contains(passport, reqFields[i]) == false {
        invalid = true
        break
      }
    }
    if !invalid {
      count++
    }
  }

  fmt.Println(count)
}
