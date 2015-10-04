package main

import (
	"fmt"
	"math"
	"unicode/utf8"
	"unsafe" //using Sizeof
)

func usingString() {
	var s1 string = "한글"
	var s2 string = "hello"
	var s3 string = "한글"
	var s4 string = "Go"

	fmt.Println("-- unit9.3 문자열 수정하기")
	var s5 string = "Hello, World\n"
	s5 = "abcdefg"

	fmt.Println(s5[0])
	// s5[0] = 'z'

	fmt.Println("-- unit9.2 문자열 배열")
	fmt.Printf("%c\n", s2[1])

	fmt.Println("-- unit9.2 문자열 연산하기")
	fmt.Println(s1 == s3)
	fmt.Println(s1 + s4 + s3)
	fmt.Println("안녕하세요" + s1)

	fmt.Println("-- unit9.1 문자열 길이 구하기")
	fmt.Println(len(s1))
	fmt.Println(len(s2))
	fmt.Println(utf8.RuneCountInString(s1))
}

func usingBool() {
	var num1 int = 3
	var num2 int = 10

	fmt.Println("-- 불 비교 --")
	fmt.Println(num1 > num2)
	fmt.Println(num1 < num2)
	fmt.Println(num1 != num2)
	fmt.Println(num1 >= num2)
	fmt.Println(num1 <= num2)

	var b1 bool = true
	var b2 bool = false

	fmt.Println("-- && --")
	fmt.Println(b1 && b1)
	fmt.Println(b1 && b2)
	fmt.Println(b2 && b2)
	fmt.Println("-- || --")
	fmt.Println(b1 || b1)
	fmt.Println(b1 || b2)
	fmt.Println(b2 || b2)
	fmt.Println("-- ! --")
	fmt.Println(!b1)
	fmt.Println(!b2)
}

func usingConstant() {
	const age int = 10
	const name string = "Maria"
	// const score int //compile error 초기화 안했음
	const score int = 90

	// age = 20 //compile error const 변경 불가
	// name = "Grace" //compile error const 변경 불가
	const age1 = 10       //대입하는 자료형으로 결정
	const name2 = "Maria" //대입하는 자료형으로 결정
	// const addr //compile error
}
func main() {
	fmt.Println("-- unit11 Constant 사용하기")
	usingConstant()

	fmt.Println("-- unit10 Bool 사용하기")
	usingBool()

	fmt.Println("-- unit9 뮨자열 사용하기")
	usingString()

	fmt.Println("-- unit8-8 변수의 크기 구하기")
	var num11 int8 = 1
	var num12 int16 = 1
	var num13 int32 = 1
	var num14 int64 = 1

	fmt.Println(unsafe.Sizeof(num11))
	fmt.Println(unsafe.Sizeof(num12))
	fmt.Println(unsafe.Sizeof(num13))
	fmt.Println(unsafe.Sizeof(num14))

	fmt.Println("-- unit8-7 overflow, underflow")

	var num7 uint8 = math.MaxUint8
	var num8 uint16 = math.MaxUint16
	var num9 uint32 = math.MaxUint32
	var num10 uint64 = math.MaxUint64

	/*	Overflow 컴파일 오류 발생
		var num7 uint8 = math.MaxUint8 + 1
		var num8 uint16 = math.MaxUint16 + 1
		var num9 uint32 = math.MaxUint32 + 1
		var num10 uint64 = math.MaxUint64 + 1
	*/

	/*
		var num7 uint8 = 0 - 1
		var num8 uint16 = 0 - 1
		var num9 uint32 = 0 - 1
		var num10 uint64 = 0 - 1
	*/

	fmt.Println(num7)
	fmt.Println(num8)
	fmt.Println(num9)
	fmt.Println(num10)

	fmt.Println("--------------------")
	var num3 int = 3
	var num4 float32 = 2.2

	// fmt.Println(num3 + num4)
	fmt.Println(float32(num3) + num4)
	fmt.Println(num3 + int(num4))

	fmt.Println("--------------------")

	var num5 int16 = 10
	var num6 int32 = 80000

	fmt.Println(num5 + int16(num6))

	fmt.Println("--------------------")
	fmt.Println("-- unit8-6 숫자 연산하기")
	var num1 uint8 = 3
	var num2 uint8 = 2

	fmt.Println(num1 + num2)
	fmt.Println(num1 - num2)
	fmt.Println(num1 * num2)
	fmt.Println(num1 / num2)
	fmt.Println(num1 % num2)
	fmt.Println(num1 << num2)
	fmt.Println(num1 >> num2)
	fmt.Println(^num1)

	fmt.Println("--------------------")

	i := 10

	if i >= 5 {
		fmt.Println("5 이상")
	}

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	fmt.Println("안녕 ~ 고랑~ unit6~")
}
