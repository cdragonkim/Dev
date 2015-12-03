package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	fmt.Println(strings.Contains("Hello, world!", "wo"))
	fmt.Println(strings.ContainsAny("Hello, world", "w o"))
	fmt.Println(strings.Count("Hello Helium", "He"))

	s1 := []string{"Hello,", "world"}
	fmt.Println(strings.Join(s1, " "))

	s2 := strings.Split("Hello, world", " ")
	fmt.Println(s2[1])

	s3 := strings.Fields("Hello, world")
	fmt.Println(s3[1])

	f := func(r rune) bool {
		return unicode.Is(unicode.Hangul, r)
	}

	s4 := strings.FieldsFunc("Hello안녕Hello", f)
	fmt.Println(s4)

	fmt.Println(strings.Repeat("Hello", 10))
	fmt.Println(strings.Replace("Hello, world", "world", "go", 1))
	fmt.Println(strings.Replace("Hello Hello", "llo", "Go", 2))
}
