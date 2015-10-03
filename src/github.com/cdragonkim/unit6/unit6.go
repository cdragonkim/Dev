package main

import (
	"fmt"
)

func main() {
	fmt.Println("안녕 ~ 고랑~ unit6~")
	i := 10

	if i >= 5 {
		fmt.Println("5 이상")
	}

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	fmt.Println("--------------------")
}
