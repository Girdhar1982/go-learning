package main

import (
	"fmt"
	"net/http"
)
const portNumber=":8080"
// each file must have main
func main() {
http.HandleFunc("/",Home)
http.HandleFunc("/about",About)
http.ListenAndServe(portNumber,nil)
}

func Home(w http.ResponseWriter,r *http.Request){
	n,err := fmt.Fprintf(w,"Home Page!")
	if err!= nil {fmt.Println(err)}
	fmt.Printf("bytes Written %d\n",n)
}

func About(w http.ResponseWriter,r *http.Request){
	sum,_:=addValues(5,2)
	_,_ = fmt.Fprintf(w,fmt.Sprintf("About Page!! %d",sum) )
}

func addValues(x,y int) (int, error){
	sum:= x + y
	return sum,nil
}