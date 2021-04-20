package main

import "fmt"

/*

Operadores racionales

>  (mayor)
<  (menor)
>= (mayor o igual)
<= (menor o igual)
== (igual)
!= (distinto)


Operadores matematicos

+ (más)
- (menos)
* (producto)
/ (división)
% (resto de una división)  Ej.:  x=13%5;  {se guarda 3}

*/

func main() {
	var num1, num2 int
	fmt.Print("Ingrese el primer valor:")
	fmt.Scan(&num1)
	fmt.Print("Ingrese el segundo valor:")
	fmt.Scan(&num2)

	//TODO: No estas obligado en usar los parentesis enla condición
	if num1 > num2 {
		fmt.Print("El mayor es ", num1)
	} else {
		fmt.Print("El manor es ", num2)
	}
}
