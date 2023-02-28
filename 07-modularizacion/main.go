package main

import (
	"fmt"
	"paquetes/models"
)

func main() {
	// mensajes.Hola()
	// mensajes.Imprimir()

	// cua1 := figuras.Cuadrado{Lado: 8}
	// figuras.Medidas(&cua1)

	// cir1 := figuras.Circulo{Radio: 5}
	// figuras.Medidas(&cir1)

	p1 := models.Persona{}
	p1.Constructor("Juan", 25)

	fmt.Println(p1)
	fmt.Println(p1.GetNombre())

	p1.SetNombre("David")
	fmt.Println(p1)

}
