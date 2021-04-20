package main

import "fmt"

func main() {
	var edades [5]int
	var nombres [5]string
	for f := 0; f < len(edades); f++ {
		fmt.Print("Ingrese nombre:")
		fmt.Scan(&nombres[f])
		fmt.Print("Ingrese edad:")
		fmt.Scan(&edades[f])
	}
	fmt.Println("Personas mayores de edad.")
	for f := 0; f < len(edades); f++ {
		if edades[f] >= 18 {
			fmt.Print(nombres[f], " ")
		}
	}
}
