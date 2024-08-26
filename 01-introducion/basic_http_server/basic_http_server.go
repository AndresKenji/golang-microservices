package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	port := 8080

	http.HandleFunc("/holamundo", holaMundoHandler)
	
	log.Println("Iniciando el servidor en el puerto:", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v",port),nil))
}

func holaMundoHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w,"Hola Mundo")
}