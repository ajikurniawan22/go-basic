package main

import "fmt"

func main() {
	//static type data
	var x int
	x = 10
	fmt.Println(x)
	fmt.Printf("x memiliki tipe data %T\n", x)

	//dynamic variable
	z := "aji kurniawan"
	fmt.Println(z)

	a := 10
	b := 20
	fmt.Println(a + b)
}
