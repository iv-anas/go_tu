package main

import (
    "fmt"
    "math"
)

const gl string = " const and var can be declare outside main function"

func main(){
	fmt.Println(gl)

	const rad= 7

	const d=math.Pi

	fmt.Println("area of circle:",float64(d*rad*rad))
}

// output
// const and var can be declare outside main function
// area of circle: 153.93804002589988