package dynamicFuncPractices

import (
	"fmt"
	"reflect"
)

func h(args []reflect.Value) []reflect.Value {
	// 매개변수와 리턴값은 반드시 []reflect.Value 사용
	fmt.Println("Hello, world!")
	return nil
}

func DaynamicFuncPrac11() {
	var hello func() // 함수 담을 변수 선언

	fn := reflect.ValueOf(&hello).Elem()
	// hello 주소 넘김 뒤 Elem으로 값 정보

	v := reflect.MakeFunc(fn.Type(), h) // h함수 정보 생성

	fn.Set(v)
	// hello 값 정보 fn에 h의 함수 정보 v를 설정하여 함수 연결

	hello()
}
