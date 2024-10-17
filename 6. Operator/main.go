package main

import "fmt"

func main() {
	// Operator Aritmatika
	var valAritmatika = (((2+6)%3)*4 - 2) / 3

	// Operator Perbandingan
	var isEqual = (valAritmatika == 2)

	fmt.Printf("nilai %d (%t) \n", valAritmatika, isEqual)

	// Operator Logika
	left := false
	right := true

	leftAndRight := left && right
	fmt.Printf("left && right \t(%t) \n", leftAndRight)

	leftOrRight := left || right
	fmt.Printf("left || right \t(%t) \n", leftOrRight)

	rightReverse := !right
	fmt.Printf("!right \t(%t) \n", rightReverse)

	// \t digunakan untuk indent tabulasi atau bisa dibilang merapikan tampilan output

}
