package main

import "fmt"

func main() {
	sum := (100 * (100 + 1)) / 2
	fmt.Println(sum)
}

/*
// ให้เขียนโปรแกรมหาค่าผลรวมตั้งแต่ 1-100 โดยให้คำนึงถึงประสิทธิภาพให้มากที่สุด

func main() {
    n := 100
    sum := 0

    for i := 1; i <= n; i++ {
        sum += i
    }

    fmt.Println("Sum:", sum)
}

*/
