package main

import (
	"fmt"
	"github.com/cdragonkim/unit41/printFuncPractices"
)

func main() {
	var i int

	fmt.Scanf("%d", &i)

	switch i {
	case 1:
		fmt.Println("-- 출력 함수 테스트 1 --")
		printFuncPractices.PrintFuncPrac11()
		fmt.Println("***********************")
	case 2:
		fmt.Println("-- 출력 함수 테스트 2 --")
		printFuncPractices.PrintFuncPrac12()
		fmt.Println("***********************")
	}

}
