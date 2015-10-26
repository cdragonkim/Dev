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
func NewRectangle(width, height int) *Rectangle {
	return &Rectangle{width, height}
}

func StructPractice2() {
	rect := NewRectangle(20, 10)
	fmt.Println(rect)
}

func rectangleArea(rect *Rectangle) int {
	return rect.width * rect.height
}

func StructPractice3() {
	rect := Rectangle{30, 30}
	area := rectangleArea(&rect)

	fmt.Println(area)
}

func rectangleScaleA(rect *Rectangle, factor int) {
	rect.width = rect.width * factor
	rect.height = rect.height * factor
}

func rectangleScaleB(rect Rectangle, factor int) {
	rect.width = rect.width * factor
	rect.height = rect.height * factor
}

func StructPractice4() {
	rect1 := Rectangle{30, 30}
	rectangleScaleA(&rect1, 10)
	fmt.Println(rect1)

	rect2 := Rectangle{30, 30}
	rectangleScaleB(rect2, 10)
	fmt.Println(rect2)
}

//----------------------------------------------

func (rect *Rectangle) area() int {
	return rect.width * rect.height
}

func StructPractice5() {
	rect := Rectangle{10, 20}
	fmt.Println(rect.area())
}

//----------------------------------------------

func (rect *Rectangle) scaleA(factor int) {
	rect.width = rect.width * factor
	rect.height = rect.height * factor
}

func (rect Rectangle) scaleB(factor int) {
	rect.width = rect.width * factor
	rect.height = rect.height * factor
}

func StructPractice6() {
	rect1 := Rectangle{30, 30}
	rect1.scaleA(10)
	fmt.Println(rect1)

	rect2 := Rectangle{30, 30}
	rect2.scaleB(10)
	fmt.Println(rect2)
}

//----------------------------------------------

func (_ Rectangle) dummy() {
	fmt.Println("dummy")
}

func StructPractice7() {
	var r Rectangle
	r.dummy()
}
