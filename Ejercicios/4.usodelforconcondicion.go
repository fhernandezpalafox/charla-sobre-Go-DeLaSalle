package main

import "fmt"

func main() {

	x := 1
	for x <= 100 {
		fmt.Print(x, "-")
		x = x + 1
	}

	// Ejemplo de código
	var n int
	fmt.Println("Programa 2")
	fmt.Println("Ingrese el valor final:")
	fmt.Scan(&n)
	y := 1
	for y <= n {
		fmt.Print(y, "-")
		y = y + 1
	}

}
