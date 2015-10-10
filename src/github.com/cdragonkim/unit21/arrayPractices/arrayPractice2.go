package arrayPractices

import (
	"fmt"
)

func ArrayPractice24() {
	a := [5]int{32, 43, 55, 33, 76}

	for _, value := range a {
		fmt.Println(value)
	}
}
func ArrayPractice23() {
	a := [5]int{32, 43, 55, 33, 76}

	for value := range a {
		fmt.Println(value)
	}
}
func ArrayPractice22() {
	a := [5]int{32, 43, 55, 33, 76}

	for i, value := range a {
		fmt.Println(i, value)
	}
}
func ArrayPractice21() {
	a := [5]int{32, 43, 55, 33, 76}

	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}
}
