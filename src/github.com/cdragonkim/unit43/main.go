package main

import (
	"fmt"
	"github.com/cdragonkim/unit43/sPrintScanPractices"
)

// func Sprint(a …interface{}) string: 값을 그대로 문자열로 만듦
// func Sprintln(a …interface{}) string: 값을 그대로 문자열로 만든 뒤 문자열 끝에 개행 문자(\n)를 붙임
// func Sprintf(format string, a …interface{}) string: 형식을 지정하여 문자열을 만듦
// func Sscan(str string, a …interface{}) (n int, err error): 공백, 개행 문자로 구분된 문자열에서 입력을 받음
// func Sscanln(str string, a …interface{}) (n int, err error): 공백으로 구분된 문자열에서 입력을 받음
// func Sscanf(str string, format string, a …interface{}) (n int, err error): 문자열에서 형식을 지정하여 입력을 받음
func main() {
	var i int
	fmt.Scanf("%d", &i)

	switch i {
	case 1:
		fmt.Println("-- Sprint ln f Test 1 --")
		sPrintScanPractices.SPrintScanPrac11()
		fmt.Println("***********************")
	case 2:
		fmt.Println("-- Sprint ln f Test 1 --")
		sPrintScanPractices.SPrintScanPrac12()
		fmt.Println("***********************")
	}
}
