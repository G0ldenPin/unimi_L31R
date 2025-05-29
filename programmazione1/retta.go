package main

import (
	"fmt"
)

func main() {
	var m, q, x, y int
	fmt.Print("Inserisci m e q: ")
	fmt.Scan(&m, &q)
	fmt.Print("Inserisci x e y: ")
	fmt.Scan(&x, &y)

	if ((m * x) + q) == y {
		fmt.Println("Il punto appartiene alla retta")
	} else if ((m * x) + q) < y {
		fmt.Println("Il punto sta sopra alla retta")
	} else {
		fmt.Println("Il punto sta sotto alla retta")
	}
}
