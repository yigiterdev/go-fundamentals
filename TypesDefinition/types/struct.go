package types

import "fmt"

type User struct {
	Id   int
	Name string
}

func (u User) PrettyPrint() {
	fmt.Println("User id:", u.Id, "name:", u.Name)
}

var u1 User = User{1, "Ahmet"}
