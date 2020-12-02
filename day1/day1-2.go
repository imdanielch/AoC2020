// what is the product of the three entries that sum to 2020?

package main

import (
  "fmt"
  "io/ioutil"
  "strconv"
  "strings"
)

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

  // Assign cap to avoid resize on every append
  nums := make([]int, 0, len(lines))


  for _, line := range lines {
    if len(line) == 0 { continue }

    n, err := strconv.Atoi(line)
    check(err)
    nums = append(nums, n)
  }

  // iterate over int slice, get the remainder, continue the iteration while checking if the number is smaller than the remainder, then do it one more time for the 3rd number.
  for i1, num1 := range nums {
    remainder := 2020 - num1
    for i2, num2 := range nums[i1:] {
      if num2 < remainder {
        remainder := remainder - num2
        for _, num3 := range nums[i1 + i2:] {
          if remainder - num3 == 0 {
            fmt.Println(num1 * num2 * num3)
            break
          }
        }
      }
    }
  }

}
