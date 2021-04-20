package main

import "fmt"

func main() {
	var valor int
	suma := 0
	for {
		fmt.Print("Ingrese un valor (0 para finalizar):")
		fmt.Scan(&valor)
		if valor == 0 {
			break
		}
		suma = suma + valor
	}
	fmt.Print("La suma de todos los valores ingresados es:", suma)
}
