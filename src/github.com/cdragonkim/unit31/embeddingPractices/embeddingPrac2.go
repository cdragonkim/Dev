package embeddingPractices

import (
	"fmt"
)

type PersonA struct {
	name string
	age  int
}

func (p *PersonA) greeting() {
	fmt.Println("Hello~")
}

type StudentA struct {
	Person
	school string
	grage  int
}

func EmbeddingPrac21() {
	var s StudentA
	s.Person.greeting()
	s.greeting()
}
