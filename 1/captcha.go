package main

import "fmt"
import "io/ioutil"
import "log"
import "os"
import "strconv"
import "bytes"

func captcha(b []byte) int{
  sum := 0

  for i, v := range b {
    if i > 0 && b[i-1] == v {
      number, _ := strconv.Atoi(string(v))
      sum += number
    }
  }

  //Compare first and last digigts
  if b[0] == b[len(b)-1] {
    number, _ := strconv.Atoi(string(b[0]))
    sum += number
  }

  return sum
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

  //Trim newline from inputfile
  buf = bytes.Trim(buf, "\x0A")

  fmt.Println(captcha(buf))
}
