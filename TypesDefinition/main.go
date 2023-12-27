package main

import (
	"fmt"

	"yigiterdev/types-definitions/types"
)

func main() {
	var u1 types.User = types.User{Id: 1, Name: "Ahmet"}

	u1.PrettyPrint()

	max := types.Instructor{
		Id:        1,
		FirstName: "Max",
		LastName:  "Verstappen",
		Score:     100,
	}

	kyle := types.NewInstructor("Kyle", "Larson")

	fmt.Println(max.FullName())
	fmt.Println(kyle.FullName())
}
