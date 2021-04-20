package main

import "fmt"

func main() {
	var sueldos [5]int //Como se declara el arreglo
	for f := 0; f < 5; f++ {
		fmt.Print("Ingrese valor del sueldo:")
		fmt.Scan(&sueldos[f])
	}
	fmt.Println("Listado de sueldos.")
	fmt.Println(sueldos)
}
