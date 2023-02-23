package main

import "fmt"

func main() {

	//Función anonima
	/*
		func() {
			fmt.Println("Hola desde la función anonima")
		}()*/

	myfunc := func(nombre string) string {
		return fmt.Sprintf("Hola, %s desde la función anonima", nombre)
	}

	fmt.Println(myfunc("Juan David"))

}
