package main

import (
	"fmt"
)

func main() {
	x := 1

	for {
		x++

		if x == 5 {
			continue
		}
		fmt.Printf("hello %d\n", x)

		if x == 10 {
			break
		}
	}
}
