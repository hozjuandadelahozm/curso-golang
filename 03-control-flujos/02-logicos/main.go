package main

import (
	"fmt"
)

func main() {

	//Not
	//fmt.Println(!false)

	//And
	//fmt.Println(true && false)

	//Or
	//fmt.Println(true || false)

	b := 2

	r := b == 2 || !(4 > b)

	fmt.Println(r)
}
