package main

import "fmt"

func saludar(nombre string) {
	fmt.Println("Hola,", nombre)
}

func sumar(num1, num2 int) int {
	return num1 + num2
}

func main() {
	saludar("Juan David")
	resultado := sumar(10, 40)
	fmt.Println("La suma es:", resultado)
}
