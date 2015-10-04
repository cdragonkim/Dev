package core02

import (
	"fmt"
)

func OpBit1() {
	var g uint8 = 3
	h := ^g

	fmt.Printf("****%08b\n", g)
	fmt.Printf("^   %08b\n\n", h)

	f := 112
	f >>= 3
	fmt.Printf(">>= %08b\n\n", f)

	e := 7
	e <<= 2
	fmt.Printf("<<= %08b\n\n", e)

	c := 112
	d := c >> 3
	fmt.Printf("*****%08b\n", c)
	fmt.Printf(">> 3 %08b\n\n", d)

	a := 7
	b := a << 2
	fmt.Printf("*****%08b\n", a)
	fmt.Printf("<< 2 %08b\n\n", b)

}

func OpBit2() {
	a := 3
	b := -2
	c := +a
	d := +b
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println("  ")
}

func OpBit3() {
	a := 3
	b := -2
	c := -a
	d := -b
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println("  ")
}
