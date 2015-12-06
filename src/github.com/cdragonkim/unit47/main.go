package main

import "fmt"
import "strconv"

func main() {
	var i int
	fmt.Scanf("%d", &i)

	switch i {
	case 1:
		var num1 int
		var err error

		num1, err = strconv.Atoi("100")
		fmt.Println(num1, err)

		num1, err = strconv.Atoi("10t")
		fmt.Println(num1, err)

		var s string
		s = strconv.Itoa(100)
		fmt.Println(s)

		s = strconv.FormatBool(true)
		fmt.Println(s)

		s = strconv.FormatFloat(1.2, 'f', -1, 32)
		fmt.Println(s)

		s = strconv.FormatInt(-10, 10)
		fmt.Println(s)

		s = strconv.FormatUint(32, 16)
		fmt.Println(s)
	case 2:
		var s []byte = make([]byte, 0)
		s = strconv.AppendBool(s, true)
		fmt.Println(string(s))

		s = strconv.AppendFloat(s, 1.3, 'f', -1, 32)
		fmt.Println(string(s))

		s = strconv.AppendInt(s, -10, 10)
		fmt.Println(string(s))

		s = strconv.AppendUint(s, 32, 16)
		fmt.Println(string(s))
	case 3:
		var err1 error
		var b1 bool

		b1, err1 = strconv.ParseBool("true")
		fmt.Println(b1, err1)

		var num11 float64
		num11, err1 = strconv.ParseFloat("1.3", 64)
		fmt.Println(num11, err1)

		var num12 int64
		num12, err1 = strconv.ParseInt("-10", 10, 32)
		fmt.Println(num12, err1)

		var num13 uint64
		num13, err1 = strconv.ParseUint("20", 16, 32)
		fmt.Println(num13, err1)

	}

}
