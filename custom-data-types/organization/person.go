package organization

import (
	"errors"
	"fmt"
	"strings"
)

type Name struct {
	first string
	last  string
}

func (n Name) FullName() string {
	return fmt.Sprintf("%s %s", n.first, n.last)
}

type Handler struct {
	handle string
	name   string
}

// TwitterHandler
// a type TwitterHandler = Handler // type alias - it copies the fields of an object as well as the method sets.
// a type TwitterHandler Handler // type declaration - it copies the fields of an object to the new type
type TwitterHandler string

func (th TwitterHandler) RedirectUrl() string {
	cleanHandler := strings.TrimPrefix(string(th), "@")
	return fmt.Sprintf("https://www.twitter.com/%s", cleanHandler)
}

type Person struct {
	Name
	twitterHandler TwitterHandler
	Citizen
	Conflict
}

func NewPerson(firstName, lastName string, c Citizen) Person {
	return Person{
		Name: Name{
			first: firstName,
			last:  lastName,
		},
		Citizen: c,
	}
}

func (p *Person) Id() string {
	return fmt.Sprintf("%s", p.Citizen.Id())
}

func (p *Person) SetTwitterHandler(handler TwitterHandler) error {
	if len(handler) == 0 {
		p.twitterHandler = handler
	} else if !strings.HasPrefix(string(handler), "@") {
		return errors.New("twitter handler must start with an @ symbol")
	}

	p.twitterHandler = handler
	return nil
}

func (p *Person) TwitterHandler() TwitterHandler {
	return p.twitterHandler
}
