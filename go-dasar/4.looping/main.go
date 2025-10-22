package main

import "fmt"

func main() {
	// looping using for
	for i := 0; i < 5; i++ {
		fmt.Println("angka", i)
	}

	// looping using for only condition (just like while)
	var while = 0
	for while < 5 {
		fmt.Println("Angka", while)
		while++
	}

	// looping using for range
	var xs = "123" // string
	for i, v := range xs {
		fmt.Println("index = ", i, "value=", v)
	}

	for _, v := range xs {
		println("angka", v)
	}
}
