package main

import "fmt"

func main() {

	type user struct {
		Id        int
		FirstName string
		LastName  string
	}

	var u user
	u.Id = 1
	u.FirstName = "bob"
	u.LastName = "X"
	fmt.Println(u)

	u2 := user{
		Id:        1,
		FirstName: "Alice",
		LastName:  "X",
	}
	fmt.Println(u2)
}
