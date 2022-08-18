package main

import (
	"fmt"
	"log"
	"time"
)

//each file must have main
func main() {
//fmt.Println(varibaleAndFunctions()) ; value1,value2:=functionreturnTwo(); fmt.Println(value1,value2)
//mainForPointer()
mainForTypesAndStructs()
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
func mainForPointer(){
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

//Types and Structs
var s ="seven"; // this variable will be available Globally in this file

type User struct{ //Capital U defines that it is public and will be visible outside the package
FirstName string
LastName string
Phone string
Age int
BirthDate time.Time
}


func mainForTypesAndStructs(){
s2:="six"
log.Println("global s variable",s)
fmt.Println("value of local s2 variable -" , s2)
first,second := changeValue("XXXX")
log.Println("global s variable",s,first,second)


user:= User{
	FirstName: "Manish",
	LastName: "Girdhar",
}
log.Println(user.FirstName,user.LastName,user.Phone)
}

func changeValue(s string)(string, string){
log.Println("s is input of this function ",s )
return "s is changed ?? ", s
}