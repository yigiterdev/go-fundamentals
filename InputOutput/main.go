package main

import "fmt"

// global variable
const url = "http://google.com"

func main() {
	printData()
	fmt.Println(add(1, 2))
	fmt.Println(addAndSubstract(1, 2))
	stateTax, cityTax := calculateTax(100)
	fmt.Println(stateTax)
	fmt.Println(cityTax)
}
