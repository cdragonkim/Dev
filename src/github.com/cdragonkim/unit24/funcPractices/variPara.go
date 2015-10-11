package funcPractices

import (
	"fmt"
)

func sum3(n ...int) int {
	total := 0
	for _, value := range n {
		total += value
	}
	return total
}

func FuncPara1() {
	r := sum3(1, 2, 3, 4, 5)
	fmt.Println(r)

	n := []int{1, 2, 3, 4, 5}
	r1 := sum3(n...)
	fmt.Println(r1)
}
