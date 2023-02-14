package main

import "fmt"

func main() {

	if nombre, edad := "Juan", 25; nombre == "Juand" {
		fmt.Println("Hola,", nombre)
	} else {
		fmt.Printf("Nombre: %s -  Edad: %d\n", nombre, edad)
	}

	//Obtener valor de mapa
	users := make(map[string]string)

	users["Juan"] = "juan@gmail.com"
	users["David"] = "david@gmail.com"

	//correo, ok := users["Juan"]

	if correo, ok := users["David"]; ok {
		fmt.Println(correo)
	} else {
		fmt.Println("No fue posible obtener el valor")
	}

}
