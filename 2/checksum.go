package main

import "fmt"
import "io/ioutil"
import "log"
import "os"
import "sort"
import "strconv"
import "bytes"

func checksum(b []byte) int{
  result := 0

  for _, row := range bytes.Split(b, []byte("\n")) {
    fields := bytes.Fields(row)
    var ints []int
    for _, field := range fields {
      number, _ := strconv.Atoi(string(field))
      ints = append(ints, number)
    }
    sort.Ints(ints)
    result += ints[len(ints)-1] - ints[0]
  }
  return result
}

func main() {
  if len(os.Args) != 2 {
    fmt.Println("Usage: takes a file of integers as only argument.")
    os.Exit(1)
  }

  buf, err := ioutil.ReadFile(os.Args[1])
  if err != nil {
    log.Fatal(err)
  }

  fmt.Println(checksum(buf))
}
