package helpers

import "time"

type User struct { //Capital U defines that it is public and will be visible outside the package
	FirstName string
	LastName  string
	Phone     string
	Age       int
	BirthDate time.Time
}