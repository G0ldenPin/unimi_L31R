package main

import (
	"fmt"
)

func main() {
	fmt.Print("Inserisci un numero:")
	var n int
	fmt.Scan(&n)
	if n%10 == 0 {
		fmt.Println(n, " è multiplo di 10")
	} else {
		fmt.Println(n, "non è multiplo di 10")
	}
}
