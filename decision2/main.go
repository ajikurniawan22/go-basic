package main

import (
	"fmt"
)

func main() {
	x := 90

	//expresion switch
	switch x {
	case 60:
		fmt.Println("Nilai C")
	case 80:
		fmt.Println("Nilai B")
	case 90:
		fmt.Println("Nilai A")
	default:
		fmt.Println("Nilai Tidak terpenuhi")
	}

	switch {
	case x == 60:
		fmt.Println("Nilai C")
	case x == 80:
		fmt.Println("Nilai B")
	case x == 90:
		fmt.Println("Nilai A")
	default:
		fmt.Println("Nilai Tidak terpenuhi")
	}

	var y interface{}
	y = "ajikurniawan"

	switch y.(type) {
	case int:
		fmt.Println("type data integer")
	case string:
		fmt.Println("Typpe data string")
	case float64:
		fmt.Println("type data float")
	default:
		fmt.Println("type data tidak di ketahui")
	}
}
