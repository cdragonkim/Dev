package funcPractices

import (
	"fmt"
)

func sum(a int, b int) int {
	return a + b
}

func diff(a int, b int) int {
	return a - b
}

func sum2(a int, b int) (r int) {
	r = a + b
	return
}

func SumAndDiff(a int, b int) (int, int) {
	return a + b, a - b
}

func FuncPractice1() {
	r := sum(1, 2)
	fmt.Println(r)

	r2 := sum2(3, 5)
	fmt.Println(r2)
}

func FuncPractice2() {
	sum, diff := SumAndDiff(6, 2)
	fmt.Println(sum, diff)
}

func FuncPractice3() {
	_, diff := SumAndDiff(6, 2)
	fmt.Println(diff)
}

func hello() (int, int, int, int, int) {
	return 1, 2, 3, 4, 5
}

func FuncPractice4() {
	a, _, c, _, e := hello()
	fmt.Println(a, c, e)
}

func SumAndDiff2(a int, b int) (sum int, diff int) {
	sum = a + b
	diff = a - b
	return
}

func FuncPractice5() {
	sum, diff := SumAndDiff2(7, 3)
	fmt.Println(sum, diff)
}

func FuncPractice6() {
	var hello1 func(a int, b int) int = sum
	world1 := sum

	fmt.Println(hello1(1, 2))
	fmt.Println(world1(2, 3))

	f := []func(int, int) int{sum, diff}
	fmt.Println(f[0](1, 2))
	fmt.Println(f[1](1, 2))
}

func FuncPractice7() {

	f := map[string]func(int, int) int{
		"sum":  sum,
		"diff": diff,
	}
	fmt.Println(f["sum"](1, 2))
	fmt.Println(f["diff"](1, 2))
}
