package main
import "fmt"
type rect struct {
    width, height int
}


func (r *rect) area() int {
    return r.width * r.height
}

func (r rect) perim() int {
    return 2*r.width + 2*r.height
}

func main(){
	var a, b int
	fmt.Scanln(&a)
	fmt.Scanln(&b)
	fmt.Println(b)
	r := rect{width: a, height: b}
	fmt.Println(r)
	fmt.Println("area: ", r.area())
    fmt.Println("perim:", r.perim())

}