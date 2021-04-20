//Sin Gorutimes
/*package main

import "fmt"

func mostrar0() {
    for f := 1; f <= 1000; f++ {
        fmt.Print("0-")
    }
}

func mostrar1() {
    for f := 1; f <= 1000; f++ {
        fmt.Print("1-")
    }
}

func main() {
    mostrar0()
    mostrar1()
}*/

/*
 En el lenguaje Go se ha creado el concepto de
 Goroutine para poder ejecutar funciones en forma concurrente,
 es decir que comience la ejecución de la función pero continúe con la ejecución
de la función main o la función desde donde se llamó la Goroutine.
*/

//Con Gorutines

package main

import "fmt"

func mostrar0() {
	for f := 1; f <= 1000; f++ {
		fmt.Print("0-")
	}
}

func mostrar1() {
	for f := 1; f <= 1000; f++ {
		fmt.Print("1-")
	}
}

func main() {
	go mostrar0()
	go mostrar1()
	var continua string
	fmt.Scan(&continua)
}
