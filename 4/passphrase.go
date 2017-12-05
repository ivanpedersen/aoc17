package main

import "fmt"
import "io/ioutil"
import "log"
import "os"
import "bytes"
import "strings"
import "sort"

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func SortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}

func valid_passphrase(b []byte) bool {
	fields := strings.Fields(string(b[:]))
	m := make(map[string]int)
	for _, field := range fields {
		m[field]++
		if m[field] > 1 {
			return false
		}
	}
	return true
}

func valid_passphrase_anagram(b []byte) bool {
	fields := strings.Fields(string(b[:]))
	m := make(map[string]int)
	for _, field := range fields {
		s := SortString(field)
		m[s]++
		if m[s] > 1 {
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

func passphrase_anagram(b []byte) int {
	result := 0
	for _, row := range bytes.Split(b, []byte("\n")) {
		if valid_passphrase_anagram(row) {
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
	fmt.Println("Part 2:", passphrase_anagram(buf))
}
