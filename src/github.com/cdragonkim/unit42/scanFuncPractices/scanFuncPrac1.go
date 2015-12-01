package scanFuncPractices

import (
	"fmt"
)

func ScanFuncPrac11() {
	var s1, s2 string
	n, _ := fmt.Scan(&s1, &s2)
	//n, _ := fmt.Scanln(&s1, &s2)

	fmt.Println("Number of Input : ", n)
	fmt.Println(s1, s2)
}

func ScanFuncPrac12() {
	var num1, num2 int
	n, _ := fmt.Scanf("%d,%d", &num1, &num2)

	fmt.Println("Number of Input : ", n)
	fmt.Println(num1, num2)
}
