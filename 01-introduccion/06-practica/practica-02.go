package main

import "fmt"

func cal_cociente(num1, num2 int) int {

	return num1 / num2
}

func cal_residuo(num1, num2 int) int {

	return num1 % num2
}

func main() {
	fmt.Println("Ejercicio #2")
	fmt.Println("==========================")
	var num1, num2, cociente, residuo int

	fmt.Print("Ingrese numero 1: ")
	fmt.Scanln(&num1)

	fmt.Print("Ingrese numero 2: ")
	fmt.Scanln(&num2)

	cociente = cal_cociente(num1, num2)
	fmt.Println("El cociente es:", cociente)

	residuo = cal_residuo(num1, num2)
	fmt.Println("El residuo es:", residuo)
	fmt.Println("==========================")
}
