// == 같다, != 같지 않다
package core02

import (
	"fmt"
)

func OpCompare1() {
	e := 1
	var p1 *int = &e
	var p2 *int = &e
	fmt.Println("포인터 비교")
	fmt.Println(p1 == p2)
	fmt.Println()

	d := map[string]int{"Hello": 1}
	fmt.Println("맵 비교")
	fmt.Println(d == nil)
	fmt.Println()

	fmt.Println(1 == 1)
	fmt.Println(3.5 == 3.5)
	fmt.Println("Hello" == "Hello")

	a := [3]int{1, 2, 3}
	b := [3]int{1, 2, 3}
	fmt.Println("배열")
	fmt.Println(a == b)
	fmt.Println()

	c := []int{1, 2, 3}
	fmt.Println("슬라이스")
	fmt.Println(c == nil)
	// 슬라리스를 nil과 비교하여
	// 메모리가 할당되었는지 확인

}
