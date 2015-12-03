package sPrintScanPractices

import (
	"fmt"
)

func SPrintScanPrac11() {
	var s1 string
	s1 = fmt.Sprint(1, 1.1, "Hello, World") //값을 문자열로
	fmt.Print(s1)

	var s2 string
	s2 = fmt.Sprintln(1, 1.1, "Hello, World")
	fmt.Print(s2)

	var s3 string
	s3 = fmt.Sprintf("%d %f %s\n", 1, 1.1, "Hello, World")
	fmt.Print(s3)

}

func SPrintScanPrac12() {
	var num1 int
	var num2 float32
	var s string

	input1 := "1\n1.1\nHello"
	n, _ := fmt.Sscan(input1, &num1, &num2, &s)
	fmt.Println("입력 개수 : ", n)
	fmt.Println(num1, num2, s)

	input2 := "1 1.1 Hello"
	n, _ = fmt.Sscanln(input2, &num1, &num2, &s)
	fmt.Println("입력 개개 : ", n)
	fmt.Println(num1, num2, s)

	input3 := "1,1.1,Hello"
	n, _ = fmt.Sscanf(input3, "%d,%f,%s", &num1, &num2, &s)
	fmt.Println("입력 갯수 : ", n)
	fmt.Println(num1, num2, s)
}
