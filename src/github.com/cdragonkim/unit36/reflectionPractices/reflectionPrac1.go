// 리플렉샨 사용하기
package reflectionPractices

import (
	"fmt"
	"reflect"
)

// ******* 변수와 구조체 type 표시 연습 ******

type Data struct {
	a, b int
}

func ReflectionPrac11() {
	var num int = 1
	fmt.Println(reflect.TypeOf(num)) // 자료형 이름

	var s string = "Hello, World!"
	fmt.Println(reflect.TypeOf(s))

	var f float32 = 1.3
	fmt.Println(reflect.TypeOf(f))

	var data Data = Data{1, 3}
	fmt.Println(reflect.TypeOf(data))
}

func ReflectionPrac12() {
	var f float64 = 1.3
	t := reflect.TypeOf(f)
	v := reflect.ValueOf(f)

	fmt.Println(t.Name())
	fmt.Println(t.Size())
	fmt.Println(t.Kind() == reflect.Float64)
	fmt.Println(t.Kind() == reflect.Int64)
	fmt.Println("------------------------")
	fmt.Println("-----값이 담긴 변수의 정보---")
	fmt.Println(v.Type())
	fmt.Println(v.Kind() == reflect.Float64)
	fmt.Println(v.Kind() == reflect.Int64)
	fmt.Println(v) // 이것도 됨 ?
	fmt.Println(v.Float())

}
