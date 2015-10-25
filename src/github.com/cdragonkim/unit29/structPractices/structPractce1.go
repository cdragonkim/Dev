package structPractices

import (
	"fmt"
)

type Rectangle struct {
	width  int
	height int
}

func StructPractice1() {
	//var rect Rectangle = Rectangle{10, 20}
	var rect1 *Rectangle

	rect1 = new(Rectangle)
	rect2 := Rectangle{45, 62}
	rect3 := Rectangle{width: 30, height: 15}

	fmt.Println(rect1)
	fmt.Println(rect2)
	fmt.Println(rect3)
}
