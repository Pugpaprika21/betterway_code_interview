package main

import "fmt"

func main() {
	numbers := []int{1, 5, 7, 2, 12, 6}

	var target int
	fmt.Print("Input: ")
	fmt.Scan(&target)

	seen := make(map[int]bool)

	for _, num := range numbers {
		complement := target - num
		if seen[complement] {
			fmt.Printf("True (**เนื่องจาก %d+%d=%d)\n", complement, num, target)
			return
		}
		seen[num] = true
	}

	fmt.Println("False")
}
