package arrayPractices

import (
	"fmt"
)

func SlicePractice6() {
	a := []int{1, 2, 3, 4, 5}
	b := a[0:3]

	fmt.Println(a)
	fmt.Println(b)
}

func SlicePractice7() {
	a := []int{1, 2, 3, 4, 5}

	fmt.Println(a[0:3])
	fmt.Println(a[1:3])
	fmt.Println(a[2:5])
}
func SlicePractice8() {
	a := []int{1, 2, 3, 4, 5}

	fmt.Println(a[:])
	fmt.Println(a[0:])
	fmt.Println(a[:5])
	fmt.Println(a[0:len(a)])

	fmt.Println(a[3:])
	fmt.Println(a[:3])
	fmt.Println(a[1:3])

}
func SlicePractice9() {
	a := [5]int{1, 2, 3, 4, 5}
	b := a[:2]
	b[0] = 9

	fmt.Println(a)
	fmt.Println(b)
}
func SlicePractice10() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8}
	b := a[0:6:8]

	fmt.Println(len(b), cap(b))
	fmt.Println(b)
}
