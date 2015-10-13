package main

import (
	"fmt"
)

func panicPractice1() {
	panic("Error !!")
	fmt.Println("Hello, world")
}

func f() {
	defer func() {
		s := recover()
		fmt.Println(s)
	}()

	panic("Error !!!")
}

func f2() {
	defer func() {
		s := recover()
		fmt.Println(s)
	}()

	a := [...]int{1, 2}

	for i := 0; i < 5; i++ {
		fmt.Println(a[i])
	}
}

func main() {
	//a := [...]int{1, 2}
	a := [...]int{1, 2, 3}
	for i := 0; i < 3; i++ {
		fmt.Println(a[i])
	}

	//panicPractice1()
	f()
	fmt.Println("Hello, world")

	f2()
	fmt.Println("Hello, world2")
}
