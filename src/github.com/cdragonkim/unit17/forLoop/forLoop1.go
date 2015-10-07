package forLoop

import (
	"fmt"
)

func ForBreakLoop() {
Loop:
	//fmt.Println("begin for loop")
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if j == 2 {
				break Loop
			}
			fmt.Println(i, j)
		}
	}
	fmt.Println(("Hello, Loop"))
}
func ForBreak1() {
	i := 0
	for {
		if i > 4 { // i가 4가 되는 순간
			break // 반복문이 중단됩니다.
		}

		fmt.Println(i)
		i = i + 1 // 변화식에서 조건을 변경합니다.
	}
}
func ForLoop1() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	fmt.Println()
	for i := 5; i > 0; i-- {
		fmt.Println(i)
	}
	fmt.Println()
	j := 0
	for j < 5 {
		fmt.Println(j)
		j = j + 1
	}
}
