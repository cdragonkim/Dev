package main

import (
	"calc" // calc 패키지 가져오기
	"fmt"
	"github.com/golang/example/stringutil"
	// or 전역패키지 . "github.com/golang/example/stringutil"
	// or 별칭 사용 strutil "github.com/golang/example/stringutil"
)

func main() {
	fmt.Println(stringutil.Reverse("olleh"))
	fmt.Println(calc.Sum(1, 2)) // calc 패키지의 Sum 함수 사용
}
