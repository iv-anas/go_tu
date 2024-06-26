package swt

import (
	"fmt"
	"time"
)

func Switch() {
	t := time.Now()

	switch{
	case t.Hour() < 12:
		fmt.Println("good morning")
	case t.Hour()>12 && t.Hour()<5:
		fmt.Println("good afternoon")
	default:
		fmt.Println("It's good evening")
	}

	

	this := func(i interface{}) {
        switch t := i.(type) {
        case bool:
            fmt.Println("I'm a bool")
        case int:
            fmt.Println("I'm an int")
        default:
            fmt.Printf("Don't know type %T\n", t)
        }
    }
	this(true)
    this(1)
    this("hey")

}


// output
// good morning
// I'm a bool
// I'm an int
// Don't know type string


