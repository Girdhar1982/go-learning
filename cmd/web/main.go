package main

import (
	"fmt"
	"net/http"
	"github.com/girdhar1982/go-learning/pkg/handlers"
)

const portNumber = ":8080"

// each file must have main
func main() {
fmt.Println("Starting Application on Port: ",portNumber)
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)
	http.ListenAndServe(portNumber, nil)
}
