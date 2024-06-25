package main

import "fmt"

type person struct {
	name string
	id   int
}

func main() {
	// Structs are mutable.
	var s1 person
	fmt.Scanln(&s1.name)
	fmt.Scanln(&s1.id)
	fmt.Println(s1)

	s := person{name: "Sean", id: 50}
    fmt.Println(s.name)

	sp := s.name
    fmt.Println(sp)

	spp := &s
    fmt.Println(spp.id)
}