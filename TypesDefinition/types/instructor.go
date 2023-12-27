package types

import "fmt"

type Instructor struct {
	Id        int
	FirstName string
	LastName  string
	Score     int
}

func NewInstructor(firstName string, lastName string) Instructor {
	return Instructor{
		FirstName: firstName,
		LastName:  lastName,
	}
}

func (i Instructor) FullName() string {
	return fmt.Sprintf("%v %v", i.FirstName, i.LastName)
}
