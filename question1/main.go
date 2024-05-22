package main

import (
	"fmt"
)

func main() {
	numbers := []int{6, 1, -50, 200, 7, 9, -13, 20}
	max := numbers[0]
	for _, num := range numbers {
		if num > max {
			max = num
		}
	}

	//fmt.Printf("ค่ามากที่สุดในอาเรย์คือ: %d\n", max)
	fmt.Println(max)
}
