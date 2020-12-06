//    Right 1, down 1.
//    Right 3, down 1. (This is the slope you already checked.)
//    Right 5, down 1.
//    Right 7, down 1.
//    Right 1, down 2.
//
// What do you get if you multiply together the number of trees encountered on each of the listed slopes?
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
  // convert []byte to []string and split
  lines := strings.Split(string(buf), "\n")
  if len(lines[len(lines) - 1]) == 0 {
    lines = lines[:len(lines) - 1]
  }

  s1 := CountTrees(lines, 1, 1)
  s2 := CountTrees(lines, 3, 1)
  s3 := CountTrees(lines, 5, 1)
  s4 := CountTrees(lines, 7, 1)
  s5 := CountTrees(lines, 1, 2)
  fmt.Println("Product of all 5 slope runs = ", s1 * s2 * s3 * s4 * s5)

}

func CountTrees (lines []string, deltaX int, deltaY int) (result int) {
  columns := len(lines[0])
  count := 0
  iX := 0
  iY := 0
  for iY < len(lines) {
    if iX >= columns {
      iX = iX % columns
    }

    fmt.Printf("%s", string(lines[iY][iX]))
    if lines[iY][iX] == byte('#') {
      count += 1
    }
    iX += deltaX
    iY += deltaY
  }
  fmt.Printf("\nTotal count of trees = %d for DeltaX: %d, DeltaY: %d\n", count, deltaX, deltaY)
  return count
}
