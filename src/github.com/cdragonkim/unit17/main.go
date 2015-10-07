// Unit 17 for 반복문 사용하기
package main

import (
	"fmt"
	"github.com/cdragonkim/unit17/forLoop"
)

func main() {
	fmt.Println("for 변수 병렬 할당")
	forLoop.ForParaVar1()
	fmt.Println("--------------")

	fmt.Println("for continue loop")
	forLoop.ForContinue2()
	fmt.Println("--------------")

	fmt.Println("for continue")
	forLoop.ForContinue1()
	fmt.Println("--------------")

	fmt.Println("for break Loop")
	forLoop.ForBreakLoop()
	fmt.Println("--------------")

	fmt.Println("for break 연습 1")
	forLoop.ForBreak1()
	fmt.Println("--------------")

	fmt.Println("for 반복문 연습 1")
	forLoop.ForLoop1()
	fmt.Println("--------------")

}
