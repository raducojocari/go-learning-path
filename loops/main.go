package main

func main() {
	//var i int
	//for i < 5 {
	//	println(i)
	//	i++
	//	if i == 3 {
	//		continue
	//	}
	//}
	//
	//for ; i < 5; i++ {
	//	println(i)
	//}
	//
	//for { // infinite loop
	//	if i < 5 {
	//		break
	//	}
	//	println(i)
	//	i++
	//
	//}
	//
	//for { // infinite loop
	//	if i < 5 {
	//		break
	//	}
	//	println(i)
	//	i++
	//}

	//iterate over collection
	slice := []int{1, 2, 3}
	for i := 0; i < len(slice); i++ {
		println(slice[i])
	}

	for i, v := range slice {
		println(i, v)
	}

	wellKnownPorts := map[string]int{"http": 80, "https": 443}
	for _, v := range wellKnownPorts {
		println(v)
	}
}
