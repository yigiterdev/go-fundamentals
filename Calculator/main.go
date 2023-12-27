package main

import "fmt"

func main() {
	var operation string
	var number1, number2 int

	fmt.Println("Calculator GO 1.0")
	fmt.Println("-----------------")
	fmt.Println("Please enter the operation to perform: ")
	fmt.Scanf("%s", &operation)
	fmt.Println("Please enter the first number: ")
	fmt.Scanf("%d", &number1)
	fmt.Println("Please enter the second number: ")
	fmt.Scanf("%d", &number2)
	fmt.Println("The result is: ", calc(operation, number1, number2))
}

func calc(operation string, number1 int, number2 int) int {
	var result int

	switch operation {
	case "+":
		result = number1 + number2
	case "-":
		result = number1 - number2
	case "*":
		result = number1 * number2
	case "/":
		result = number1 / number2
	default:
		fmt.Println("Invalid operation")
	}

	return result
}
