package maps

import (
	"fmt"
	"maps"
)

func Maps() {
	var m1 map[string]int                // Declares a nil map
	m2 := make(map[string]int)           // Creates an empty map
	m3 := map[string]int{"a": 1, "b": 2} // Creates and initializes a map
	fmt.Println("maps",m1, m2, m3)

	v1 := m3["k1"]
    fmt.Println("v1:", v1)

	fmt.Println("len:", len(m1), len(m2), len(m3))

	delete(m3, "b")
    fmt.Println("map3:", m3)

	n := map[string]int{"foo": 1, "bar": 2}
    fmt.Println("map:", n)

	n2 := map[string]int{"foo": 1, "bar": 2}
    if maps.Equal(n, n2) {
        fmt.Println("n == n2")
    }
}	
