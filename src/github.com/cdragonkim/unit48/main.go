package main

import (
	"fmt"
	"regexp"
)

func main() {
	matched, _ := regexp.MatchString("[a-zA-Z0-9]+", "Hello 100")
	fmt.Println(matched)

	matched, _ = regexp.MatchString("[가-힣]+", "안녕하세요")
	fmt.Println(matched)

	matched, _ = regexp.MatchString("^Hello", "Go Hello, world!")
	fmt.Println(matched)

	matched, _ = regexp.MatchString("\\.[a-zA-Z]+\\([0-9]+\\)$", "Hello.SetValue(20)")
	fmt.Println(matched) // true: Hello.SetValue(20)는 .영문(숫자)이므로 true

	matched, _ = regexp.MatchString("\\.[a-zA-Z]+\\([0-9]+\\)$", "Hello.SetValue(20).x")
	fmt.Println(matched) // false: Hello.SetValue(20).x는 .영문(숫자)가 아니므로 false

}
