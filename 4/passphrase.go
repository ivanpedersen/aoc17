package main

import "fmt"
import "io/ioutil"
import "log"
import "os"
import "bytes"
import "strings"

func valid_passphrase(b []byte) bool {
	fields := strings.Fields(string(b[:]))
	m := make(map[string]int)
	for _, field := range fields {
		m[field]++
		fmt.Println("m: ", m)
		if m[field] > 1 {
			return false
		}
	}
	return true
}

func passphrase(b []byte) int {
	result := 0
	for _, row := range bytes.Split(b, []byte("\n")) {
		if valid_passphrase(row) {
			result += 1
		}
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

	fmt.Println("Part 1:", passphrase(buf))
}
