package main

import (
	"custom-data-types/organization"
	"fmt"
)

func main() {
	p := organization.NewPerson("Bob", "A", organization.NewEuId("123-456", "Germany"))
	err := p.SetTwitterHandler("@jam_wils")
	if err != nil {
		fmt.Println("Error incorrect twitter handler")
	}

	println(p.Id())
	println(p.FullName())
	println(p.Country())
	println(p.TwitterHandler().RedirectUrl())
}
