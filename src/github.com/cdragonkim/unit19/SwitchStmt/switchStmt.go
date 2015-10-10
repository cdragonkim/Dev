package SwitchStmt

import (
	"fmt"
)

func SwitchStmt1() {
	i := 3

	switch i {
	case 0:
		fmt.Println(0)
	case 1:
		fmt.Println(1)
	case 2:
		fmt.Println(2)
	case 3:
		fmt.Println(3)
	default:
		fmt.Println(-1)
	}
}
func SwitchStmt2() {
	s := "world"

	switch s {
	case "hello":
		fmt.Println("hello")
	case "world":
		fmt.Println("world")
	default:
		fmt.Println("일치 문자열 없음")
	}
}
func SwitchStmt3() {
	s := "hello"
	i := 2

	switch i {
	case 1:
		fmt.Println("1")
	case 2:
		if s == "hello" {
			fmt.Println("hello 2")
			break
		}
		fmt.Println(2)
	}
}
func SwitchStmt4() {
	i := 3

	switch i {
	case 4:
		fmt.Println("4 이상")
		fallthrough
	case 3:
		fmt.Println("3 이상")
		fallthrough
	case 2:
		fmt.Println("2 이상")
		fallthrough
	case 1:
		fmt.Println("1 이상")
		fallthrough
	case 0:
		fmt.Println("0 이상")
	}
}

func SwitchStmt5() {
	i := 3

	switch i {
	case 2, 4, 6:
		fmt.Println("짝수")
	case 1, 3, 5:
		fmt.Println("홀수")
	}
}
