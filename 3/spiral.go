package main

import "fmt"
import "log"
import "os"
import "strconv"
import "math"

func square_size(n int) int {
	size := int(math.Ceil(math.Sqrt(float64(n))))
	if size%2 == 0 {
		size += 1
	}

	return size
}

func spiral_part1(number int) int {
	size := square_size(number)
	if size == 1 {
		return 0
	}
	layers := (size - 1) / 2
	cycle := number - int((size-2)*(size-2))
	innerOffset := cycle % (size - 1)

	return layers + int(math.Abs(float64(innerOffset)-float64(layers)))
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: takes an integer as only argument.")
		os.Exit(1)
	}
	arg := os.Args[1]
	number, err := strconv.Atoi(arg)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Part 1:", spiral_part1(number))
}
