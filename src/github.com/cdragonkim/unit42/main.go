package main

import (
	"fmt"
	"github.com/cdragonkim/unit42/scanFuncPractices"
)

func main() {
	var i int
	fmt.Scanf("%d", &i)

	switch i {
	case 1:
		fmt.Println("-- Scan Func Test 1 --")
		scanFuncPractices.ScanFuncPrac11()
		fmt.Println("***********************")
	case 2:
		fmt.Println("-- Scanf Func Test  --")
		scanFuncPractices.ScanFuncPrac12()
		fmt.Println("***********************")
	}
}
