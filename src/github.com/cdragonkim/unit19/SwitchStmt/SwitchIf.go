package SwitchStmt

import (
	"fmt"
	"math/rand"
	"time"
)

func SwitchIf1() {
	i := 7

	switch {
	case i >= 5 && i < 10:
		fmt.Println("5 이상, 10미만")
	case i >= 0 && i < 5:
		fmt.Println("0 이상, 5미만")
	}
}
func SwitchIf2() {
	rand.Seed(time.Now().UnixNano())
	switch i := rand.Intn(10) {
	case i >= 3 && i < 6:
		fmt.Println("3 이상, 6미만")
	case i ==9:
		fmt.Println("9")
	default:
		fmt.Println(i)
	}
}
