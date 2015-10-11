package mapPractices

import (
	"fmt"
)

func MapPractice1() {
	var a map[string]int = make(map[string]int)
	var b = make(map[string]int)
	c := make(map[string]int)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
func MapPractice2() {
	// map에 키와 값 할당
	a := map[string]int{"Hello": 10, "world": 20}
	b := map[string]int{
		"Hello": 30,
		"world": 40,
	}

	fmt.Println(a["Hello"])
	fmt.Println(b["world"])
}
func MapPractice3() {
	solarSystem := make(map[string]float32)

	solarSystem["Mercury"] = 87.969 // 맵[키] = 값
	solarSystem["Venus"] = 224.70069
	solarSystem["Earth"] = 365.25641
	solarSystem["Mars"] = 686.9600
	solarSystem["Jupiter"] = 4333.2867
	solarSystem["Saturn"] = 10756.1995
	solarSystem["Uranus"] = 30707.4896
	solarSystem["Neptune"] = 60223.3528

	fmt.Println(solarSystem["Mars"])
	fmt.Println("-------------------")
	value, ok := solarSystem["Pluto"]
	fmt.Println(value, ok)
	fmt.Println("-------------------")

	if value, ok := solarSystem["Saturn"]; ok {
		fmt.Println(ok, value)
		fmt.Println("-------------------")
	}

	for key, vlaue := range solarSystem {
		fmt.Println(key, vlaue)
	}
	for _, vlaue := range solarSystem {
		fmt.Println(vlaue)
	}
}
