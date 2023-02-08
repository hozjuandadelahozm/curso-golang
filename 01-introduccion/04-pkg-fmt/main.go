package main

import "fmt"

func main() {
	hola := "Hola"
	mundo := "Mundo"

	fmt.Println(hola)
	fmt.Println(mundo)

	nombre := "Juan David"
	edad := 25

	fmt.Printf("Hola %s edad %d \n", nombre, edad)

	fmt.Printf("Hola %v edad %v \n", nombre, edad)

	mensaje := fmt.Sprintf("Hola %s edad %d", nombre, edad)

	fmt.Println(mensaje)

	fmt.Printf("nombre: %T \n", nombre)

	fmt.Print("Ingrese otro nombre: ")
	fmt.Scanln(&nombre)

	fmt.Println("Otro nombre: ", nombre)
}
