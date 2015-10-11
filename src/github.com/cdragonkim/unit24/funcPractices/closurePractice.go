package funcPractices

import (
	"fmt"
)

func clac() func(x int) int {
	a, b := 3, 5
	return func(x int) int {
		return a*x + b
	}

}

func ClosurePractice1() {
	f := clac()

	fmt.Println(f(1))
	fmt.Println(f(2))
	fmt.Println(f(3))
	fmt.Println(f(4))
	fmt.Println(f(5))

}
