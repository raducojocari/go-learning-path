package main

type User struct {
	ID        int
	FirstName string
	LastName  string
}

func main() {
	u1 := User{
		ID:        1,
		FirstName: "Bob",
		LastName:  "Smith",
	}

	u2 := User{
		ID:        2,
		FirstName: "Ford",
		LastName:  "Prefect",
	}

	if u1.ID == u2.ID {
		println("same user!")
	} else if u1.FirstName == u2.LastName {
		println("similar user")
	} else {
		println("Difference user")
	}

}
