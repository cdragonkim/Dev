package embeddingPractices

import (
	"fmt"
)

type Person struct {
	name string
	age  int
}

func (p *Person) greeting() {
	fmt.Println("Hello~")
}

type Student struct {
	p      Person
	school string
	grage  int
}

func EmbeddingPrac11() {
	var s Student
	s.p.greeting()
}
