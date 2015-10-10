package arrayPractices

import (
	"fmt"
)

func ArrayPractice4() {
	a := [3][3]int{
		{32, 43, 55},
		{29, 22, 33},
		{33, 56, 76},
	}

	fmt.Println(a)

}
func ArrayPractice3() {
	a := [5]int{32, 43, 55, 22, 33}

	b := [5]int{
		22,
		12,
		56,
		67,
		99, // 여러 줄로 나열할 때는 마지막 요소에 콤마를 붙임
	}

	c := [...]int{23, 34, 21, 54, 32}

	d := [...]string{"Maria", "Andrew", "Jessi"}

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}
func ArrayPractice2() {
	var a [5]int = [5]int{21, 23, 43, 22, 12}
	var b = [5]int{21, 23, 43, 22, 12}
	c := [5]int{66, 44, 45, 32, 34}

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
func ArrayPractice1() {
	var a [5]int

	a[2] = 7

	fmt.Println(a)
}
