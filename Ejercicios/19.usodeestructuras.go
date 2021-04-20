package main

import "fmt"

type Producto struct {
	codigo      int
	descripcion string
	precio      float64
}

func cargar(producto *Producto) {
	fmt.Print("Ingrese el código:")
	fmt.Scan(&producto.codigo)
	fmt.Print("Ingrese la descripción:")
	fmt.Scan(&producto.descripcion)
	fmt.Print("Ingrese el precio")
	fmt.Scan(&producto.precio)
}

func imprimir(producto Producto) {
	fmt.Println("Codigo:", producto.codigo)
	fmt.Println("Descripcion:", producto.descripcion)
	fmt.Println("Precio:", producto.precio)
}

func main() {
	var producto1 Producto
	cargar(&producto1)
	imprimir(producto1)
}
