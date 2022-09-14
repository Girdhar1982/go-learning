package helpers

import (
	"math/rand"
	"time"
)

type User struct { //Capital U defines that it is public and will be visible outside the package
	FirstName string
	LastName  string
	Phone     string
	Age       int
	BirthDate time.Time
}

func RandomNumber(n int) int {
	rand.Seed(time.Now().UnixNano())
	value := rand.Intn(n)
	return value
}
