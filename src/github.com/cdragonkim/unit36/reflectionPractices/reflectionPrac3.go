package reflectionPractices

import (
	"fmt"
	"reflect"
)

func ReflectionPrac31() {
	var a *int = new(int)
	*a = 1

	fmt.Println("*a TypeOf : ", reflect.TypeOf(a))
	fmt.Println("*a ValueOf : ", reflect.ValueOf(a))
	//fmt.Println(reflect.ValueOf(a).Int())
	fmt.Println("*a Elem : ", reflect.ValueOf(a).Elem())
	fmt.Println("*a Elem Int ", reflect.ValueOf(a).Elem().Int())
	fmt.Println("---------------")

	var b interface{}
	b = 1

	fmt.Println("I/F b TypeOf : ", reflect.TypeOf(b))
	fmt.Println("I/F b ValueOf : ", reflect.ValueOf(b))
	fmt.Println("I/F b ValueOf : ", reflect.ValueOf(b).Int())
	//fmt.Println(reflect.ValueOf(b).Elem())
}
