package main

import (
	"fmt"
)

// 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610, 987, 1597, 2584, 4181

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}

	prev := 0
	curr := 1

	for i := 2; i <= n; i++ {
		temp := curr
		curr = prev + curr
		prev = temp
	}

	return curr
}

func main() {
	var input int
	fmt.Print("Input: ")
	fmt.Scan(&input)

	result := fibonacci(input)
	fmt.Printf("Output: %d\n", result)
}
