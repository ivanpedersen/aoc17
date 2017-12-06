package main

import "fmt"
import "io/ioutil"
import "log"
import "os"
import "bytes"
import "strconv"

func get_high_bank(list []int) int {
	index := 0
	value := 0

	for i := 0; i < len(list); i++ {
		if list[i] > value {
			value = list[i]
			index = i
		}
	}
	return index
}

func redistribute(list []int, index int) []int {
	bank_value := list[index]
	list[index] = 0
	position := index + 1
	for {
		if bank_value > 0 {
			if position >= len(list) {
				position = 0
			}
			list[position]++
			bank_value--
			position++
		} else {
			break
		}
	}
	return list
}

func mem(list []int) (int, []int) {
	steps := 0
	position := 0
	states := make(map[string]int)
	for {
		position = get_high_bank(list)
		list = redistribute(list, position)
		steps++
		state := fmt.Sprint(list)
		states[state]++
		if states[state] > 1 {
			break
		}
	}
	return steps, list
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: takes a file of integers separated with '\\t' as only argument.")
		os.Exit(1)
	}

	buf, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	list := []int{}
	for _, row := range bytes.Split(buf, []byte("\t")) {
		i, _ := strconv.Atoi(string(row))
		list = append(list, i)
	}

	result, list2 := mem(list)
	result2, _ := mem(list2)
	fmt.Println("Part 1:", result)
	fmt.Println("Part 2:", result2-1)
}
