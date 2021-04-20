package main

import "fmt"

func main() {
	//Para inicializar uso de la función make
	paises := make(map[string]int)
	paises["argentina"] = 40000000
	paises["españa"] = 46000000
	paises["brasil"] = 190000000
	paises["uruguay"] = 3400000

	//Sin el uso de Make

	/*

		 paises := map[string]int{
	        "argentina" : 40000000,
	        "españa"    : 46000000,
	        "brasil"    : 190000000,
	        "uruguay"   : 3400000,
	    }

	*/
	fmt.Println(paises)
	fmt.Println("La cantidad de habitantes de españa es:", paises["españa"])
}
