package main

//TODO: Se importa lo que se va usar si no se usa mandara error de igual manera las variables
import "fmt"

func main() {
	var lado int
	var superficie int

	fmt.Print("Ingrese el valor del lado del cuadrado:")
	fmt.Scan(&lado)
	superficie = lado * lado
	fmt.Print("La superficie del cuadrado es:", superficie)
}
