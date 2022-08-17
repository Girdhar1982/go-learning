package main

import (
	"fmt"
	"log"
)

//each file must have main
func main() {
//fmt.Println(varibaleAndFunctions()) ; value1,value2:=functionreturnTwo(); fmt.Println(value1,value2)
variablesForPointer()
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

//Pointers
func variablesForPointer(){
var myString string
myString="Green"
log.Println("myString is Set to," ,myString)
fmt.Println("Let's change the value using Pointers.")
changeUsingPointer(&myString)
log.Println("myString is now Set to," ,myString)
}

func changeUsingPointer(s *string){ //s is pointer of string
	log.Println("Address of String ," ,s)
newValue:="Red"; *s =newValue;
}