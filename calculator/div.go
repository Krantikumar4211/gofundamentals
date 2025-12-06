package main

func Div(a int, b int) int {
	if b == 0 {
		return 0 // avoid division by zero crash
	}
	return a / b
}
