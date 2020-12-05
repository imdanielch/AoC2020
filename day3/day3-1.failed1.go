package main

import (
  "fmt"
  "io"
  "os"
  "bytes"
)

// helper function to check for errors
func check(e error) {
  if e != nil {
    panic (e)
  }
}

func main() {
  file, err := os.Open("./input")
  check(err)
  lines, columns, err := CountMatrix(file)
  check(err)
  file.Close()
  fmt.Printf("lines: %d, columns: %d\n", lines, columns)
  iX := 0
  iY := 0

}

func CountMatrix (r io.Reader) (int, int, error) {
  buf := make([]byte, 32*1024)
  count := 0
  columns := 0
  lineSep := []byte("\n")
  lineSep2 := byte('\n')

  for {
    c, err := r.Read(buf)
    if columns < 1 {
      columns += bytes.IndexByte(buf[:c], lineSep2)
    }
    count += bytes.Count(buf[:c], lineSep)

    switch {
    case err == io.EOF:
      return count, columns, nil

    case err != nil:
      return count, columns, err
    }
  }
}
