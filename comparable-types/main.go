package main

import "fmt"

type Name struct {
	First string
	Last  string
}

type OtherName struct {
	First string
	Last  string
}

type Animal interface {
	Name() string
}

type Mamal interface {
	Animal
	Eat(food string) string
}

type dog string
type cat string

func (d dog) Name() string {
	return string(d)
}

func (d dog) Eat(food string) string {
	return fmt.Sprintf("%s eats %s", string(d), food)
}

func (d cat) Name() string {
	return string(d)
}

func (d cat) Eat(food string) string {
	return fmt.Sprintf("%s eats %s", string(d), food)
}

func NewDog(value string) Mamal {
	return dog(value)
}

func NewCat(value string) Mamal {
	return cat(value)
}

//Used with equals method

type YetAnotherName struct {
	First  string
	Last   string
	Middle []string
}

func (n YetAnotherName) Equals(otherName YetAnotherName) bool {
	return n.First == otherName.First &&
		n.Last == otherName.Last && len(n.Middle) == len(otherName.Middle)
}

///

func main() {
	name1 := Name{First: "Bob", Last: "B"}
	name2 := Name{First: "Bob", Last: "B"}
	//name3 := OtherName{First: "Bob", Last: "B"}

	d := NewDog("Toto")
	c := NewCat("Kitty")
	fmt.Printf("%T\n", d)
	fmt.Printf("%T\n", c)

	if name2 == name1 {
		println("it matches")
	}

	if d == c {
		println("it matches")
	}

	// zero value comparison
	name4 := Name{First: "", Last: ""}
	name5 := &name4
	name5 = nil
	if name4 == (Name{}) {
		println("It's empty")
	}

	if name5 == nil {
		println("Comparing with nil matches as well but it's more expensive since it puts pointers on the stack")
	}
}
