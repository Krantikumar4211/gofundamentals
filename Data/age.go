package main

import "fmt"

func Age() {
	var age1 int
	var age2 int
	var age3 int

	fmt.Println("Enter 1st Member Age")
	fmt.Scanln(&age1)

	fmt.Println("Enter 2st Member Age")
	fmt.Scanln(&age2)

	fmt.Println("Enter 3st Member Age")
	fmt.Scanln(&age3)

	fmt.Println("1st Member Age :", age1)
	fmt.Println("2st Member Age :", age2)
	fmt.Println("3st Member Age :", age3)
}
func main() {

	Age()
}
