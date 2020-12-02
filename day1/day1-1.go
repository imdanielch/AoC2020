// Find the two entries that sum to 2020; what do you get if you multiply them together?

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
  // read file into buffer. returns a []byte
  data, err := ioutil.ReadFile("./input")
  check(err)

  // convert []byte to []string and split
  lines := strings.Split(string(data), "\n")

  // Assign cap to avoid resize on every append.
  nums := make([]int, 0, len(lines))


  for _, line := range lines {
    if len(line) == 0 { continue }

    n, err := strconv.Atoi(line)
    check(err)
    nums = append(nums, n)
  }

  for _, num := range nums {
    counterpart := 2020 - num
    if Contains(nums, counterpart) {
      fmt.Println(counterpart * num)
      break
    }
  }


}

// checks if int slice contains the integer in question
func Contains(slice []int, val int) (bool) {
  for _, entry := range slice {
    if entry == val {
      return true
    }
  }
  return false
}
