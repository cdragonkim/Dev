// 슬라이스 사용하기
package arrayPractices

import (
	"fmt"
)

func SliceCopyPractice2() {
	a := []int{1, 2, 3, 4, 5}
	fmt.Println(len(a), cap(a))
	a = append(a, 6, 7)
	fmt.Println(len(a), cap(a))
}
func SliceCopyPractice1() {
	a := []int{1, 2, 3, 4, 5}
	b := make([]int, 3)

	copy(b, a)
	b[0] = 9
	fmt.Println(a)
	fmt.Println(b)
}
func SlicePractice5() {
	a := []int{1, 2, 3}
	var b []int

	b = a
	b[0] = 9
	fmt.Println(a)
	fmt.Println(b)
}

func SlicePractice4() {
	a := [3]int{1, 2, 3}
	var b [3]int

	b = a
	b[0] = 9
	fmt.Println(a)
	fmt.Println(b)
}

func SlicePractice3() {
	a := []int{1, 2, 3}
	b := []int{4, 5, 6}

	a = append(a, b...)
	fmt.Println(a)
}

func SlicePractice2() {
	a := []int{1, 2, 3}
	a = append(a, 4, 5, 6)
	fmt.Println(a)
}

func SlicePractice1() {
	var a []int = make([]int, 5)
	var b = make([]int, 5)
	c := make([]int, 5)
	d := []int{23, 34, 99, 56, 57}
	e := []int{
		1,
		2,
		3,
		4,
		5,
	}
	var f = make([]int, 5, 10) //길이,용량
	g := make([]int, 5, 10)

	fmt.Println(g[4])
	fmt.Println(g[5])
	fmt.Println(g[6])
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Printf("길이 : %d\n", len(f))
	fmt.Printf("용량 : %d\n", cap(f))
}
