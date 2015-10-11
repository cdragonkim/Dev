package funcPractices

import (
	"fmt"
)

func FuncPractice8() {
	func() {
		fmt.Println("Hello, world")
	}()

	func(s string) {
		fmt.Println(s)
	}("Hello, world")

	r := func(a int, b int) int {
		return a + b
	}(1, 2)

	fmt.Println(r)
}
