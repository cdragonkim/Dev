package SwitchStmt

import (
	"fmt"
)

func SwitchStmt1() {
	i := 3

	switch i {
	case 0:
		fmt.Println(0)
	case 1:
		fmt.Println(1)
	case 2:
		fmt.Println(2)
	case 3:
		fmt.Println(3)
	default:
		fmt.Println(-1)
	}
}
