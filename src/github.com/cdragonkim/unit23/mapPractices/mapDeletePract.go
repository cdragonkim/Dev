package mapPractices

import (
	"fmt"
)

func MapDeletePractice1() {
	a := map[string]int{"hello": 10, "world": 20}
	delete(a, "world")
	fmt.Println(a)
}
