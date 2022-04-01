package main

import "fmt"

func main() {
	//array
	var arr [3]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3

	fmt.Println(arr)
	fmt.Println(arr[2])

	arr2 := [3]int{1, 2, 2}
	fmt.Println(arr2)

	//slice
	slice := arr[:] //create a slice from arr from begining to the end
	arr[1] = 42
	slice[2] = 27
	fmt.Println(slice)

	slice2 := []int{1, 2, 3}
	fmt.Println(slice2)

	slice2 = append(slice2, 4, 100, 103)
	fmt.Println(slice2)

	slice3 := slice2[4:]
	fmt.Println(slice3)

	slice4 := slice2[:2]
	fmt.Println(slice4)

	slice5 := slice2[1:2]
	fmt.Println(slice5)

	//maps
	m := map[string]int{"foo": 42}
	fmt.Println(m)
	fmt.Println(m["foo"])

	delete(m, "foo")
}
