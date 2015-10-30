package reflectionPractices

import (
	"fmt"
	"reflect"
)

func ReflectionPracX1() {
	var x float64 = 3.4
	v := reflect.ValueOf(x)

	fmt.Println("type : ", v.Type())
	fmt.Println("Kind is float64 ? ", v.Kind() == reflect.Float64)
	//fmt.Println("Kind id float64 ? ", v.Type() == reflect.Float64)
	fmt.Println("Value : ", v.Float())
}
