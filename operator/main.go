package main

import (
	"fmt"
)

func main() {
	x := 10
	y := 5

	//test operator
	zAdd := x + y
	fmt.Println(zAdd)

	//test relational
	i := 10
	j := 10
	fmt.Println(i == j)
	fmt.Println(i != j)
}
