package main

import "fmt"

//Suma de dos numeros

func suma(num1, num2 int) int {
	return num1 + num2
}

func main() {
	fmt.Println("Ejercicio #1")
	fmt.Println("==========================")
	var num1, num2, resultado int

	fmt.Print("Ingrese numero 1: ")
	fmt.Scanln(&num1)

	fmt.Print("Ingrese numero 2: ")
	fmt.Scanln(&num2)

	resultado = suma(num1, num2)
	fmt.Println("El resultado es:", resultado)
	fmt.Println("==========================")
}
