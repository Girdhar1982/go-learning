package main

import (
	"encoding/json"
	"fmt"
	"github.com/girdhar1982/go-learning/helpers"
	"log"
)

// each file must have main
func main() {
	//fmt.Println(varibaleAndFunctions()) ; value1,value2:=functionreturnTwo(); fmt.Println(value1,value2)
	//mainForPointer()
	//mainForTypesAndStructs()
	//mainForStructWithFunctions()
	//mainForMapAndSlices()
	//	mainForDecisionOperators()
	//mainForLooping()
	//mainForInterface()
	//mainForChannels()
	mainForJsonReadWrite()
}

// Variables and Functions
func varibaleAndFunctions() string {
	var whatToSay string
	var i int
	i = 42
	who := "Everyone"
	whatToSay = "Thank you .."
	return (whatToSay + who + " I am " + string(i) + " Years Old")
}
func functionreturnTwo() (string, string) {
	var whatToSay string
	who := "Everyone"
	whatToSay = "Thank you .."
	return (whatToSay + who), "Good bye."
}

// Pointers
func mainForPointer() {
	var myString string
	myString = "Green"
	log.Println("myString is Set to,", myString)
	fmt.Println("Let's change the value using Pointers.")
	changeUsingPointer(&myString)
	log.Println("myString is now Set to,", myString)
}

func changeUsingPointer(s *string) { //s is pointer of string
	log.Println("Address of String ,", s)
	newValue := "Red"
	*s = newValue
}

// Types and Structs
var s = "seven" // this variable will be available Globally in this file

func mainForTypesAndStructs() {
	s2 := "six"
	log.Println("global s variable", s)
	fmt.Println("value of local s2 variable -", s2)
	first, second := changeValue("XXXX")
	log.Println("global s variable", s, first, second)

	user := helpers.User{
		FirstName: "Manish",
		LastName:  "Girdhar",
	}
	log.Println(user.FirstName, user.LastName, user.Phone)
}

func changeValue(s string) (string, string) {
	log.Println("s is input of this function ", s)
	return "s is changed ?? ", s
}

//Functions in Struct

type myStruct struct {
	FirstName string
}

// m is receiver here will tie the function to struct
func (m *myStruct) printFirstName() string {
	return m.FirstName
}

func mainForStructWithFunctions() {
	var myVar myStruct
	myVar.FirstName = "Manish"
	log.Println(myVar.printFirstName())
}

// maps and slices
func mainForMapAndSlices() {
	//must use make function to make new map
	myMap := make((map[string]string))
	myMap["MyName"] = "Manish"
	myMap["YourName"] = "Nisha"
	log.Println(myMap["MyName"], " and ", myMap["YourName"])
	userMap := make(map[string]helpers.User)
	userMap["First"] = helpers.User{FirstName: "Manish", LastName: "Girdhar"}
	log.Println(userMap["First"])
	userMap["Second"] = helpers.User{FirstName: "Nisha", LastName: "Girdhar"}
	log.Println(userMap["Second"])
	log.Println(userMap)
	//slices == Array
	var firstNames []string //example one
	firstNames = append(firstNames, "Mehul")
	firstNames = append(firstNames, "Javin")
	log.Println(firstNames)
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} //example 2
	log.Println(numbers[0:5])

	name := []string{"first", "second", "last", "middle"} //example 3
	log.Println(name)
}

// Decision Structures
func mainForDecisionOperators() {
	myName := "Nisha"
	if myName == "Manish" {
		fmt.Println("You are ", myName)
	} else if myName == "Nisha" {
		fmt.Println("Hello  ", myName)
	}
	// && = and operator
	// || = or Operator
	// == equal operator

	switch myName {
	case "Manish":
		fmt.Println("You are ", myName)
	case "Nisha":
		fmt.Println("Hello ", myName)
	default:
		fmt.Println("You are nobody here ", myName)

	}

}

//Looping

func mainForLooping() {
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}

	names := []string{"first", "second", "last", "middle"} //example 3

	for i, name := range names {
		fmt.Println(i, name)
	}

	for _, name := range names {
		fmt.Println(name)
	}

	myMap := make((map[string]string))
	myMap["MyName"] = "Manish"
	myMap["YourName"] = "Nisha"
	for key, value := range myMap {
		fmt.Println(key, value)
	}

	lines := "This is my first line"
	for i, l := range lines {
		fmt.Println(i, l)
	}

}

//interfaces

type Employee interface {
	Department() string
	Skills() string
}

type Worker struct {
	Name   string
	salary int
}

func (w *Worker) Department() string {
	return "Library"
}
func (w *Worker) Skills() string {
	return "Unknown"
}

type Labour struct {
	Name  string
	wages int
}

func (w *Labour) Department() string {
	return "Garden"
}
func (w *Labour) Skills() string {
	return "digging"
}

func mainForInterface() {
	worker := Worker{Name: "Doug", salary: 100}

	PrintInfo(&worker)
	labour := Labour{Name: "Ram", wages: 10}
	PrintInfo(&labour)
}

func PrintInfo(e Employee) {
	fmt.Println(e.Department(), e.Skills())
}

// Channels
func CalculateValue(intChan chan int) {
	const numPool = 100
	randomNumber := helpers.RandomNumber(numPool)
	intChan <- randomNumber
}
func mainForChannels() {
	intChan := make(chan int)
	defer close(intChan) //what ever comes after defer execute when current function is done (Channel is needed to be closed)

	go CalculateValue(intChan) //Go routine take Channel and run it
	num := <-intChan //assign to num whatever comes from int channel
	log.Println(num)
}

//json read & write
type Person struct {
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	HairColor string `json:"hair_color"`
	HasDog bool `json:"has_dog"`
}
func mainForJsonReadWrite(){
	myJson := `
[
    {
        "first_name": "Clark",
        "last_name": "Kent",
        "hair_color": "black",
        "has_dog": true
    },
    {
        "first_name": "Bruce",
        "last_name": "Wayne",
        "hair_color": "black",
        "has_dog": false
    }
]`
var unmarshalled []Person
err := json.Unmarshal([]byte(myJson), &unmarshalled)
if err != nil {
	log.Println("Error unmarshalling json", err)
}
log.Printf("unmarshalled: %v", unmarshalled)
// write json from a struct
var mySlice []Person

var m1 Person
m1.FirstName = "Wally"
m1.LastName = "West"
m1.HairColor = "red"
m1.HasDog = false

mySlice = append(mySlice, m1)

var m2 Person
m2.FirstName = "Diana"
m2.LastName = "Prince"
m2.HairColor = "black"
m2.HasDog = false

mySlice = append(mySlice, m2)

newJson, err := json.MarshalIndent(mySlice, "", "     ") //MarshalIndent will show the formated json
if err != nil {
	log.Println("error marshalling", err)
}

fmt.Println(string(newJson))  //String to convert slice of byte to String
}