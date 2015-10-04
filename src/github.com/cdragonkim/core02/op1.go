package core02

import (
	"fmt"
)

func OpCont1() {
	var a int = 1
	var b int = 2
	var c int = b
	const d string = "Hello, World!"

	fmt.Println(a, b, c, d)

	e := 1
	f := 3.5
	g := "안녕 고랑~"

	fmt.Println(e, f, g)

	h := 1 + 2
	i := 2 + 3
	j := h + i
	k := d + g
	fmt.Println(h, i, j, k)

	l := 3 - 2
	m := 4 - 5
	n := l - m
	fmt.Println(l, m, n)

	o := 2 * 3
	p := 9 * 21
	q := o * p
	fmt.Println(o, p, q)
}

func OpCont2() {
	a := 5 / 2
	b := 12 / 4
	c := a / b
	fmt.Println(a, b, c)
}

func OpCont3() {
	a := 5 % 2
	fmt.Println(a)
}

func OpCont4() {
	a := 5
	a += 2
	fmt.Println(a)
}

func OpCont5() {
	a := "Hello, "
	a += "World"
	fmt.Println(a)
}

func OpCont6() {
	a := 5
	a -= 2
	fmt.Println(a)
}

func OpCont7() {
	a := 5
	a *= 2
	fmt.Println(a)
}

func OpCont8() {
	a := 5
	a /= 2
	fmt.Println(a)
}

func OpCont9() {
	a := 5
	a %= 2
	fmt.Println(a)
}

func OpCont10() {
	a := 3
	b := 2
	c := a & b
	fmt.Printf("%08b & %08b = %08b\n", a, b, c)
}

func OpCont11() {
	a := 3
	b := 2
	c := a | b
	fmt.Printf("%08b | %08b = %08b\n", a, b, c)
}

func OpCont12() {
	a := 3
	b := 2
	c := a ^ b
	fmt.Printf("%08b ^ %08b = %08b\n", a, b, c)
}

func OpCont13() {
	a := 255
	b := 68
	c := a &^ b
	fmt.Printf("***%08b\n", a)
	fmt.Printf("&^ %08b\n", b)
	fmt.Println("-----------")
	fmt.Printf("=  %08b\n", c)
}

func OpCont14() {
	a := 3
	b := 2
	a &= b
	fmt.Printf("%08b\n", a)
}
func OpCont15() {
	a := 3
	b := 2
	a |= b
	fmt.Printf("%08b\n", a)
}
func OpCont16() {
	a := 3
	b := 2
	a ^= b
	fmt.Printf("%08b\n", a)
}
func OpCont17() {
	a := 255
	b := 68
	a &^= b
	fmt.Printf("%08b\n", a)
}
