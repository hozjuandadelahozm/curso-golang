package main

import "fmt"

//Variable global
var mensaje string

func funcion1() {
	mensaje = "Hola desde la función uno!"
	fmt.Println(mensaje)
}

func funcion2() {
	mensaje = "Hola desde función dos!"
	fmt.Println(mensaje)
}

func main() {

	mensaje = "Hola desde la función pricipal!"
	fmt.Println(mensaje)

	defer funcion1()
	funcion2()

	println(mensaje)
}
