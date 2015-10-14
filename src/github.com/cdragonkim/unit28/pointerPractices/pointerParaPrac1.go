package pointerPractices

import (
	"fmt"
)

func hello(n int) {
	n = 2
}

func hello2(n *int) {
	*n = 2
}

func PointerParaPrac1() {
	var n int = 1
	hello(n)
	fmt.Println(n)
}

func PointerParaPrac2() {
	var n int = 1
	hello2(&n)
	fmt.Println(n)
}
