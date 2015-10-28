// 구조체 태그 가져오기
package reflectionPractices

import (
	"fmt"
	"reflect"
)

type Person struct { //구조체 태그
	name string `tag1:"이름" tag2:"Name"`
	age  int    `tag1:"나이" tag2:"Age"`
}

func ReflectionPrac21() {
	p := Person{}
	name, ok := reflect.TypeOf(p).FieldByName("name")
	fmt.Println(ok, name.Tag.Get("tag1"), name.Tag.Get("tag2"))

	age, ok := reflect.TypeOf(p).FieldByName("age")
	fmt.Println(ok, age.Tag.Get("tag1"), age.Tag.Get("tag2"))
}
