package main

import "fmt"

func printData() {
	// function scoped variable
	var message string = "Hello from Go!"
	price := 34.4
	fmt.Println(message, price, url)
}

func add(a int, b int) int {
	return a + b
}

func addAndSubstract(a int, b int) (int, int) {
	return a + b, a - b
}

func calculateTax(price float32) (float32, float32) {
	return price * 0.08, price * 1.08
}

func calculateTaxWithName(price float32) (stateTax float32, cityTax float32) {
	return price * 0.08, price * 1.08
}

func birtday(age *int) {
	if *age > 140 {
		defer panic("Invalid age")
	}

	*age++
}
