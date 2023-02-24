package main

import "fmt"

func cal_IGV(valor, porcentaje float64) float64 {

	return (porcentaje * valor) / 100
}

func cal_precio(valor, IGV float64) float64 {

	return valor + IGV
}

func main() {
	fmt.Println("Ejercicio #3")
	fmt.Println("==========================")
	var valor, IGV, prec_venta, porcentaje float64

	porcentaje = 18

	fmt.Print("Ingrese el valor de venta del producto:")
	fmt.Scanln(&valor)

	fmt.Println("El valor de venta es:", valor)

	IGV = cal_IGV(valor, porcentaje)
	fmt.Println("El IGV es:", IGV)

	prec_venta = cal_precio(valor, IGV)
	fmt.Println("El precio de venta es:", prec_venta)
	fmt.Println("==========================")
}
