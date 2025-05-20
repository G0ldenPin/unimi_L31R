package main

import (
	"fmt"
)

func main() {
	fmt.Print("Inserisci un numero: ")
	var n int
	fmt.Scan(&n)
	switch {
	case n%3 == 0:
		fmt.Println("Fizz")
	case n%5 == 0:
		fmt.Println("Buzz")
	default:
		return
	}
}
