package structPractices

import (
	"fmt"
)

func StructPractice1() {
	type Rectangle struct {
		width  int
		height int
	}
	//or
	// type Rectangle struct {
	// 	width,height int
	// }

	var rect1 Rectangle
	var rect2 *Rectangle = new(Rectangle)

	rect1.height = 20
	rect2.height - 62

	fmt.Println(rect1)
	fmt.Println(rect2)
}
