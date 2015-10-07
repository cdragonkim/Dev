package forLoop

import (
	"fmt"
)

func ForContinue1() {
	for i := 0; i < 5; i++ {
		if i == 2 {
			continue
		}
		fmt.Println(i)
	}
}

func ForContinue2() {
Loop:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if j == 2 {
				continue Loop
			}
			fmt.Println(i, j)
		}
	}
	fmt.Println("Hello, continue loop")
}
