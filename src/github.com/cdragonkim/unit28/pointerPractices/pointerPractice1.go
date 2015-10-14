package pointerPractices

import (
	"fmt"
)

func PointerPractice1() {
	var numPtr *int
	fmt.Println(numPtr)
}

func PointerPractice2() {
	var numPtr *int = new(int)
	fmt.Println(numPtr)
}

func PointerPractice3() {
	var numPtr *int = new(int)
	*numPtr = 1
	fmt.Println(*numPtr)
}

func PointerPractice4() {
	var num int = 1
	var numPtr *int = &num
	fmt.Println(numPtr)
	fmt.Println(&num)
}

// func PointerPractice5() {
// 	var numPtr *int = new(int)

// 	numPtr++              // 컨파일 에러 : 포인터 연산 허용 안함
// 	numPtr = 0xc0820062d0 // 컴파일 에러 : 주소 직접 대입 금지

// 	fmt.Println(numPtr)
// }
