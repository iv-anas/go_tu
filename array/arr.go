package array

import "fmt"

func Array() {
	var a [5]int
	fmt.Println("empty array :",a)   //zero value

	a[4]=100
	fmt.Println("geta[4] :",a[4])

	fmt.Println(len(a))

	var b = [...]int{1, 2, 3, 4, 5}
	var c = [...]int{1,2,3,4,5,6,7,8}

	fmt.Println("len of b:",len(b), "len of c:",len(c))

	for i:=0;i<len(c);i++{
		fmt.Println("value of array in c is",c[i])
	}

	for index, value:=range c{
		fmt.Println(index, value)
	}
	arr1 := [3]int{1, 2, 3}
	arr2 := arr1
	arr2[0] = 10
	fmt.Println(arr1) // Output: [1 2 3]
	fmt.Println(arr2) // Output: [10 2 3]

	var matrix [2][3]int
	matrix = [2][3]int{{1, 2, 3}, {4, 5, 6}}
	for i:=0;i<2;i++{
		for j:=0;j<3;j++{
			fmt.Print(matrix[i][j]," ")
		}
		fmt.Println()
		fmt.Println("test")
	}

}

// output
// empty array : [0 0 0 0 0]
// geta[4] : 100
// 5
// len of b: 5 len of c: 8
// value of array in c is 1
// value of array in c is 2
// value of array in c is 3
// value of array in c is 4
// value of array in c is 5
// value of array in c is 6
// value of array in c is 7
// value of array in c is 8
// 0 1
// 1 2
// 2 3
// 3 4
// 4 5
// 5 6
// 6 7
// 7 8
// [1 2 3]
// [10 2 3]
// 1 2 3
// 4 5 6