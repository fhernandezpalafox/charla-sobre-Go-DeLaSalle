package main

import "fmt"

func main() {
	//TODO: Declaración de todas las variables de forma inversa
	var num1, num2, suma, producto int
	fmt.Print("Ingrese primer valor:")
	fmt.Scan(&num1)
	fmt.Print("Ingrese segundo valor:")
	fmt.Scan(&num2)
	suma = num1 + num2
	producto = num1 * num2
	fmt.Println("La suma de los dos valores es:", suma)
	fmt.Println("El producto de los dos valores es:", producto)
}
