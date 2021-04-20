package main

import "fmt"

/*

Los punteros en el lenguaje Go son esenciales:

Para que un programa sea muy eficiente.
Modificar variables de tipo int, float64, string etc. en otras funciones.
Poder requerir memoria durante la ejecución del programa
(hay muchas situaciones donde no sabemos cuanto espacio reservar)


*/

func main() {
	var valor1 int = 10
	var valor2 int = 20
	var pe *int
	pe = &valor1
	fmt.Println("Lo apuntado por pe es:", *pe)
	fmt.Println("La direccion que almacena pe es:", pe)
	pe = &valor2
	fmt.Println("Lo apuntado por pe es:", *pe)
	fmt.Println("La direccion que almacena pe es:", pe)
}
