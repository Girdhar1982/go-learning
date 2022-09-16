package main

import (
	"fmt"
	"net/http"
)

// each file must have main
func main() {
http.HandleFunc("/",func(w http.ResponseWriter,r *http.Request){
	n,err := fmt.Fprintf(w,"Hello, World!")
	if err!= nil {fmt.Println(err)}
	fmt.Printf("bytes Written %d\n",n)})
	http.ListenAndServe(":8080",nil)
}
