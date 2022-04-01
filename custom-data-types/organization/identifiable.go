package organization

type Identifiable interface {
	// Id only functions fo in here
	Id() string
}

type Citizen interface {
	Identifiable
	Country() string
}

type Conflict interface {
	Id() string
}

type socialSecurityNumber string

func NewSocialSecurityNumber(value string) Citizen {
	return socialSecurityNumber(value)
}
func (ssn socialSecurityNumber) Id() string {
	return string(ssn)
}

func (ssn socialSecurityNumber) Country() string {
	return "USA"
}

type europeadnUnionIdentifier struct {
	id      string
	country string
}

func NewEuId(id, country string) Citizen {
	return europeadnUnionIdentifier{
		id:      id,
		country: country,
	}
}

func (e europeadnUnionIdentifier) Id() string {
	return e.id
}

func (e europeadnUnionIdentifier) Country() string {
	return e.country
}
