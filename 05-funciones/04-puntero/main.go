package main

import "fmt"

func modificarValor(cadena *string) {
	fmt.Printf("%p\n", cadena)
	*cadena = "Hola desde la función"
}

func main() {

	cadena := "Hola Mundo de Go"
	fmt.Printf("%p\n", &cadena)
	fmt.Println("Antes de la función:", cadena)

	modificarValor(&cadena) //Referencia

	fmt.Println("Despues de la función:", cadena)

}
