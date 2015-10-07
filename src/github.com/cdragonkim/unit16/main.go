package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	// var b []byte
	// var err error

	// b, err = ioutil.ReadFile("./hello.txt")

	// if err == nil {
	// 	fmt.Printf("%s\n", b)
	// }

	if b, err := ioutil.ReadFile("./hello.txt"); err == nil {
		fmt.Printf("%s\n", b)
	}

	// fmt.Println(b)
	// fmt.Println(err)
}
