package main

import "fmt"

func main() {
	a := 20
	b := 10

	//Suma
	result := a + b
	fmt.Println("Sema:", result)

	//Resta
	result = a - b
	fmt.Println("Resta:", result)

	//Multiplicación
	result = a * b
	fmt.Println("Multiplicación:", result)

	//División
	result = a / b
	fmt.Println("División:", result)

	//Division Exacta
	var div float64 = 3.0 / 2.0
	fmt.Println("División Exacta:", div)

	//Modulo
	result = 3 % 2
	fmt.Println("Modulo:", result)
}
