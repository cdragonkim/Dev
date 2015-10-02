package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println("-------UNIT 6 Start---------")
	i := 10

	if i >= 5 {
		fmt.Println("5 이상")
	}

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	fmt.Println("-------UNIT 6 End  ---------")

	fmt.Println("------- 실수 처리 1 ---------")

	var a float64 = 10.0

	for i := 0; i < 10; i++ {
		a = a - 0.1
	}

	fmt.Println(a)
	fmt.Println(a == 9.0)

	fmt.Println("------- 실수 처리 2 ---------")

	a = 10.0

	for i := 0; i < 10; i++ {
		a = a - 0.1
	}

	fmt.Println(a)

	const epsilon = 1e-14
	fmt.Println(math.Abs(a-9.0) <= epsilon)
}
