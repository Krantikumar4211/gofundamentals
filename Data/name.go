package main

import "fmt"

var age1, age2, age3 int
var name1, name2, name3 string

func main() {
	Age()

	fmt.Println("1st Member Age :", age1)
	fmt.Println("2nd Member Age :", age2)
	fmt.Println("3rd Member Age :", age3)

	fmt.Println("Enter 1st member Name: ")
	fmt.Scanln(&name1)

	fmt.Println("Enter 2st member Name: ")
	fmt.Scanln(&name2)

	fmt.Println("Enter 3st member Name: ")
	fmt.Scanln(&name3)

	fmt.Println("The First Member name is ", name1, " And Age is :", age1)
	fmt.Println("The Second Member name is ", name2, " And Age is :", age2)
	fmt.Println("The Third Member name is ", name3, " And Age is :", age3)
}
