package main

import "fmt"

var appName = "Calculator" // global var

func demo() {
	var x int = 10 // explicit type
	var y = 20     // inferred type
	z := 30        // short declaration (inside function only)

	fmt.Println(x, y, z)

	// Redeclaration using :=
	a, b := 5, 10
	a, b, c := 15, 20, 25 // c is new → allowed

	fmt.Println(a, b, c)
}
