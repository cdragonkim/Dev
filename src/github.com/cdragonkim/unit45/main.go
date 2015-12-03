package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func main() {
	var s string = "한"
	var s2 string = "안녕하세요"
	var r1 rune = '한'
	fmt.Println(unicode.Is(unicode.Hangul, r1))
	fmt.Println(unicode.Is(unicode.Latin, r1))
	fmt.Println(unicode.In(r1, unicode.Latin, unicode.Han, unicode.Hangul))

	fmt.Println(len(s))
	fmt.Println(utf8.RuneLen(r1))
	fmt.Println(utf8.RuneCountInString(s2))

	b := []byte("안녕하세요")
	r, size := utf8.DecodeRune(b)
	fmt.Printf("%c %d\n", r, size)

	r, size = utf8.DecodeRune(b[3:])
	fmt.Printf("%c %d\n", r, size)

	r, size = utf8.DecodeLastRune(b)
	fmt.Printf("%c %d\n", r, size)

	r, size = utf8.DecodeLastRune(b[:len(b)-3])
	fmt.Printf("%c %d\n", r, size)

	s = "Hello, World!"
	fmt.Printf("%c\n", s[0])
	fmt.Printf("%c\n", s[len(s)-1])

	s = "안녕하세요"
	r, _ = utf8.DecodeRuneInString(s)
	fmt.Printf("%c\n", r)

	r, _ = utf8.DecodeLastRuneInString(s)
	fmt.Printf("%c\n", r)

}
