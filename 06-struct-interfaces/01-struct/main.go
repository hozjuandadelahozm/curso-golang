package main

import "fmt"

// Struct persona

type Persona struct {
	nombre string
	edad   int
}

//metodos

func (p *Persona) imprimir() {
	fmt.Printf("Nombre: %s\nEdad: %d\n", p.nombre, p.edad)
}

func (p *Persona) cambiarNombre(nuevoNombre string) {
	p.nombre = nuevoNombre
}

//Herencia

type Empleado struct {
	Persona
	sueldo float64
}

func main() {

	p1 := Persona{"Juan", 25}

	p1.imprimir()
	// fmt.Println(p1)
	// p1.nombre = "David"

	p1.cambiarNombre("David")

	p1.imprimir()
	// fmt.Println(p1)

	p2 := Persona{
		nombre: "De la hoz",
		edad:   26,
	}

	p2.imprimir()
	p2.cambiarNombre("Diego")
	p2.imprimir()

	// fmt.Println(p2)

	em1 := Empleado{
		sueldo: 500,
	}
	em1.nombre = "Pedro"
	em1.edad = 30
	em1.imprimir()
	fmt.Println(em1)

}
