package main

import "fmt"

func plus(a, b, c int) int {
	return a + b + c
}

func main() {
	var a, b, c int
	fmt.Scanln(&a)
	fmt.Scanln(&b)
	fmt.Scanln(&c)

	res:=plus(a,b,c)
	fmt.Println(res)
}