package main

import (
	"fmt"
	"github.com/cdragonkim/unit26/deferPractices"
)

func main() {
	fmt.Println("--- defer 연습 ---")
	deferPractices.DeferPractice1()
	fmt.Println()

	fmt.Println("--- defer 연습 2 ---")
	deferPractices.DeferPractice2()
	fmt.Println()

	fmt.Println("--- defer 연습 3 ---")
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i)
	}
	fmt.Println()

	fmt.Println("--- defer 연습 3 ---")
	deferPractices.ReadHello()
	fmt.Println()
}
