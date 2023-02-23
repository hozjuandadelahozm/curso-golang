package main

import "fmt"

func sumar(nombre string, numeros ...int) (string, int) {

	mensaje := fmt.Sprintf("La suma de %s es: ", nombre)
	var total int
	for _, num := range numeros {
		total += num
	}

	return mensaje, total
}

func main() {

	mensaje, result := sumar("Juan David", 10, 20, 40, 70, 60)
	fmt.Println(mensaje, result)
}
