package main

import "fmt"

func retornarSuperficie(lado int) int {
	superficie := lado * lado
	return superficie
}

func main() {
	var valor int
	fmt.Print("Ingrese el valor del lado del cuadrado:")
	fmt.Scan(&valor)
	sup := retornarSuperficie(valor)
	fmt.Print("La superficie del cuadrado es ", sup)
}
