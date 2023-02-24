package main

import (
	"fmt"
	"os"
)

func main() {

	//Función
	defer func() {
		if error := recover(); error != nil {
			fmt.Println("Ups!, al paracer el programa no finalizo de forma correcta")
		}
	}()

	if file, error := os.Open("holaj.txt"); error != nil {
		panic("No fue posible óbtener el archivo!!")
	} else {

		defer func() {
			fmt.Println("Cerrando el archivo!")
			file.Close()
		}()

		contenido := make([]byte, 254)
		long, _ := file.Read(contenido)
		contenidoArchivo := string(contenido[:long])
		fmt.Println(contenidoArchivo)
	}

}
