package main

import "fmt"

// global variable
const url = "http://google.com"

func main() {
	printData()
	stateTax, _ := calculateTaxWithName(100)
	fmt.Println(stateTax)

	age := 30
	birtday(&age)
	fmt.Println(age)
	day("Saturday")
}
