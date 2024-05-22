package main

import (
	"fmt"
)

func main() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("GoodLuck")
		} else if i%3 == 0 {
			fmt.Println("Good")
		} else if i%5 == 0 {
			fmt.Println("Luck")
		} else {
			fmt.Println(i)
		}
	}
}
