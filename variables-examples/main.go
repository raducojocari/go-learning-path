package main

import "fmt"

const applePi = 3.1415

const (
	first  = 1
	second = "second"
)

const (
	apple = iota + 6
	pear  = 2 << iota
	cherry
)

func main() {
	//primitives
	var i int
	i = 42
	fmt.Println(i)

	var f float32 = 3.14
	fmt.Println(f)

	firstName := "Bob"
	fmt.Println(firstName)
	b := true
	fmt.Println(b)

	c := complex(3, 4)
	fmt.Println(c)

	r, im := real(c), imag(c) // real and imaginary numbers :))
	fmt.Println(r, im)

	//pointers
	var name *string = new(string)

	*name = "Bob"
	fmt.Println(*name)

	//-----------------------
	middleName := "Alice"
	fmt.Println(middleName)

	ptr := &middleName
	fmt.Println(ptr, *ptr)

	middleName = "Tricia"
	fmt.Println(ptr, *ptr)
	//-------------------------

	//constants

	const pi = 3.1415
	fmt.Println(pi)

	const d int = 3
	fmt.Println(d + 3)

	fmt.Println(float32(d) + 1.2)

	//iota
	fmt.Println(applePi)

	fmt.Println(first, second)
	fmt.Println(apple, pear, cherry)

}
