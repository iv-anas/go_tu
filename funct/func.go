package funct

import "fmt"

func plus(a, b, c int) int {
	return a + b + c
}

func vals() (int, int) {
    var a,b int
	fmt.Scanln(&a)
	fmt.Scanln(&b)

	return a,b

}
func Func() {
	var a, b, c int
	fmt.Scanln(&a)
	fmt.Scanln(&b)
	fmt.Scanln(&c)

	res:=plus(a,b,c)
	fmt.Println(res)

	
	_,e:=vals()
	fmt.Println(e)
}