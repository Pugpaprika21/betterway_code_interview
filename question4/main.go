package main

import (
	"fmt"
)

func main() {
	arr1 := []string{"Apple", "Papaya", "Banana", "Orange"}
	arr2 := []string{"Mango", "Banana", "Apple", "Lemon"}

	common := make(map[string]bool)
	for _, fruit := range arr1 {
		common[fruit] = true
	}

	for _, fruit := range arr2 {
		if common[fruit] {
			fmt.Printf("%s ", fruit)
			delete(common, fruit)
		}
	}
}
