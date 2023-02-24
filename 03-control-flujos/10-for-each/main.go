package main

import "fmt"

func main() {

	nombres := [...]string{"Juan", "David", "Alex"}

	// for i := 0; i < len(nombres); i++ {
	// 	fmt.Println(nombres[i])
	// }

	for indice, elemento := range nombres {
		fmt.Println(indice, elemento)
	}

	//Imprimir solo los elementos
	for _, elemento := range nombres {
		fmt.Println(elemento)
	}

}
