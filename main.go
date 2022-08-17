
package main

import "fmt"

//each file must have main 
func main() {

fmt.Println(varibaleAndFunctions())
value1,value2:=functionreturnTwo()
fmt.Println(value1,value2)

}

// Variables and Functions
func varibaleAndFunctions() string{
	var whatToSay string
	var i int
	i=42;
	who:="Everyone"
	whatToSay = "Thank you .."
	return (whatToSay + who +  " I am "+ string(i) +" Years Old")
}
func functionreturnTwo() (string,string ){
	var whatToSay string
	who:="Everyone"
	whatToSay = "Thank you .."
	return  (whatToSay + who), "Good bye."
}