package organization

type A struct {
	LastName  string
	FirstName string
}

type B = A

type C A

func (e A) a1() {
	e.FirstName = "BOB"
}

func (h B) a2() {
	h.FirstName = "BOB"
	h.a1()
}

func (a C) a3() {
	a.LastName = "Bob"
	//a.a1() - does not work
}
