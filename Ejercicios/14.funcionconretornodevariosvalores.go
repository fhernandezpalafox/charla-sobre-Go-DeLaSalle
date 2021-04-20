package main

import "fmt"

func mayorMenor(v1, v2 int) (int, int) {
	if v1 > v2 {
		return v1, v2
	} else {
		return v2, v1
	}
}

func main() {
	var x1, x2 int
	fmt.Print("Ingrese primer valor:")
	fmt.Scan(&x1)
	fmt.Print("Ingrese segundo valor:")
	fmt.Scan(&x2)
	var may, men int
	may, men = mayorMenor(x1, x2)
	fmt.Println("El mayor es:", may)
	fmt.Println("El menor es:", men)
}
