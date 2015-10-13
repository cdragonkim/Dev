package deferPractices

import (
	"fmt"
	"os"
)

func hello() {
	fmt.Println("Hello")
}

func world() {
	fmt.Println("world")
}

func HelloWorld() {
	defer func() {
		fmt.Println("world")
	}()
	func() {
		fmt.Println("Hello")
	}()
}

func DeferPractice1() {
	defer world()
	hello()
	hello()
	hello()
}

func DeferPractice2() {
	HelloWorld()
}

func ReadHello() {
	file, err := os.Open("hello.txt")
	defer file.Close()

	if err != nil {
		fmt.Println(err)
		return
	}
	buf := make([]byte, 100)
	if _, err = file.Read(buf); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(buf))
}
