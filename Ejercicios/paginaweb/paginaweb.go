package main

import (
	"fmt"
	"net/http"
)

func inicioFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola mundo desde web")
}

func acercaFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Acerca de")
}

func main() {
	//http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("./public"))))
	//http.HandleFunc("/", inicioFunc)
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/acerca", acercaFunc)
	http.HandleFunc("/contacto", contactoHandler)
	fmt.Printf("Servidor corriendo en http://localhost:3001")
	http.ListenAndServe(":3001", nil)

}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./public/index.html")
}

func contactoHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./public/contacto.html")
}
