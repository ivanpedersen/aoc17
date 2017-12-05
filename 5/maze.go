package main

import "fmt"
import "io/ioutil"
import "log"
import "os"
import "bytes"
import "strconv"

func convert_bytearray_to_intarray(b []byte) []int {
	list := []int{}
	for _, row := range bytes.Split(b, []byte("\n")) {
		i, _ := strconv.Atoi(string(row))
		list = append(list, i)
	}
	return list
}

func maze_part1(b []byte) int {
	list := convert_bytearray_to_intarray(b)
	position := 0
	result := 0

	for {
		if position >= len(list) {
			break
		}
		jump := list[position]
		list[position]++
		position += jump

		result += 1
	}

	return result
}

func maze_part2(b []byte) int {
	list := convert_bytearray_to_intarray(b)
	position := 0
	result := 0

	for {
		if position >= len(list) {
			break
		}
		jump := list[position]
		if jump >= 3 {
			list[position]--
		} else {
			list[position]++
		}
		position += jump

		result += 1
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

	fmt.Println("Part 1:", maze_part1(buf))
	fmt.Println("Part 2:", maze_part2(buf))
}
