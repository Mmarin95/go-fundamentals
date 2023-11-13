package main

import "fmt"

func main() {
	var operation string
	var num1, num2 int

	fmt.Print("Enter first number: ")
	fmt.Scanf("%d", &num1)

	fmt.Print("Enter second number: ")
	fmt.Scanf("%d", &num2)

	fmt.Print("Enter operation (+, -, *, /): ")
	fmt.Scanf("%s", &operation)

	switch operation {
	case "+":
		fmt.Printf("%d + %d = %d", num1, num2, num1+num2)
	case "-":
		fmt.Printf("%d - %d = %d", num1, num2, num1-num2)
	case "*":
		fmt.Printf("%d * %d = %d", num1, num2, num1*num2)
	case "/":
		fmt.Printf("%d / %d = %d", num1, num2, num1/num2)
	default:
		fmt.Println("Invalid operation")
	}

	fmt.Println()
}
