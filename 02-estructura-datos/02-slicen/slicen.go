package main

import "fmt"

func main() {
	//Slicen
	numeros := []int{1, 2, 3}

	fmt.Println(numeros, len(numeros))

	//Agregar datos
	numeros = append(numeros, 4)
	numeros = append(numeros, 5)

	fmt.Println(numeros, len(numeros))

	//Sub Slicen
	subSlicen := numeros[:2]

	numeros[0] = 100

	fmt.Println(subSlicen)
	fmt.Println(numeros)
}
