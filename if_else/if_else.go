package main

import (
	"fmt"
	// "errors"
)

func main() {
	
	// if x ==  {
    //     return x, errors.New("empty name")
    // }

	n:=9

	if n%2==0{
		fmt.Println(n," n is even")
	}else{
		fmt.Println(n," n is odd")
	}

	if x:=18; x%2==0 && x%3==0{
		fmt.Println(x," divisible by 6")
	}

	j:=34

	if j<0{
		fmt.Println(j," number is negative")
	}else if j<10{
		fmt.Println(j," number is single digit positive")
	}else{
		fmt.Println(j," number is greater than 10")
	}

}

// output
// 9  n is odd
// 18  divisible by 6
// 34  number is greater than 10