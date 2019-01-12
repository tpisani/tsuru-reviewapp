package main

import (
	"fmt"
)

func main() {
	testSum([]int{2, 2, 2, 4}, 10)
	testSum([]int{-1, -2, -3, -4, 5}, -5)
}

func testSum(numbers []int, expected int) {
	sum := Sum(numbers)
	if sum != expected {
		message := fmt.Sprintf("Expected the sum of %v to be %d but instead got %d!", numbers, expected, sum)
		panic(message)
	}
}

func Sum(numbers []int) int {
	sum := 0
	// This bug is intentional
	for n := range numbers {
		sum += n
	}
	return sum
}
