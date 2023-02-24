package main

import (
	"fmt"
	"strings"
)

// Closures
func repeat(n int) func(cadena string) string {

	return func(cadena string) string {
		return strings.Repeat(cadena, n)
	}
}

func main() {

	repeat3 := repeat(3)
	fmt.Println(repeat3("Hola"))

	//Función anonima
	/*
			func() {
				fmt.Println("Hola desde la función anonima")
			}()

		myfunc := func(nombre string) string {
			return fmt.Sprintf("Hola, %s desde la función anonima", nombre)
		}

		fmt.Println(myfunc("Juan David"))
	*/

}
