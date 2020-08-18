package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var myString = "Hello Golang"
	var myStringTwo = "10"

	fmt.Println(myString)
	fmt.Println(len(myString))

	myStringUpper := strings.ToUpper(myString)
	fmt.Println(myStringUpper)

	myStringLower := strings.ToLower(myString)
	fmt.Println(myStringLower)

	resultContains := strings.Contains(myString, "Go")
	fmt.Println(resultContains)

	resultSplit := strings.Split(myString, " ") //return dynamic array / slice
	for _, v := range resultSplit {
		fmt.Println(v)
	}

	//konversi string to integer
	resultConvint, err := strconv.Atoi(myStringTwo)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resultConvint * 5)

	//konversi integer to string
	resultConvStr := strconv.Itoa(10)
	fmt.Println(resultConvStr)

}
