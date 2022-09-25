package main

import (
	"fmt"
	"net/http"
	"github.com/girdhar1982/go-learning/pkg/handlers"
)

const portNumber = ":8080"

// each file must have main
func main() {

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)
	fmt.Println("Starting Application on Port: ",portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
