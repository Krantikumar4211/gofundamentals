package main

import (
	"fmt"
)

func main() {
	for {
		fmt.Println("\n====== SIMPLE CALCULATOR ======")
		fmt.Println("1. Add")
		fmt.Println("2. Subtract")
		fmt.Println("3. Multiply")
		fmt.Println("4. Divide")
		fmt.Println("5. Exit")
		fmt.Print("Enter choice: ")

		var choice int
		fmt.Scanln(&choice)

		if choice == 5 {
			fmt.Println("Exiting... Goodbye!")
			break
		}

		var a, b int
		fmt.Print("Enter first number: ")
		fmt.Scanln(&a)

		fmt.Print("Enter second number: ")
		fmt.Scanln(&b)

		switch choice {
		case 1:
			fmt.Println("Result:", Add(a, b))
		case 2:
			fmt.Println("Result:", Sub(a, b))
		case 3:
			fmt.Println("Result:", Mul(a, b))
		case 4:
			fmt.Println("Result:", Div(a, b))
		default:
			fmt.Println("Invalid choice!")
		}
	}
}
