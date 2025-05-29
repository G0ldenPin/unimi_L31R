package main

import (
	"fmt"
)

/*
Se a , b e c sono numeri naturali e a² + b² = c² , si dice che la terna di numeri a ,
b e c è una terna pitagorica.
*/

func main() {
	var soglia int
	fmt.Print("Inserisci la soglia: ")
	fmt.Scan(&soglia)
	if soglia <= 0 {
		fmt.Println("BRUTTO CRETINO NON FUNZIONA SE METTI UNA SOGLIA NEGATIVA")
	} else {
		ternepitagoriche(soglia)
	}
}

func ternapitagoricatrue(a, b, c int) bool {
	x := (a * a) + (b * b)
	y := c * c

	return x == y
}

func ternepitagoriche(soglia int) {
	for a := 1; a < soglia-1; a++ {
		for b := 1; b < soglia-1; b++ {
			for c := b; c < soglia-1; c++ {
				if ternapitagoricatrue(a, b, c) {
					fmt.Println("(", a, ",", b, ",", c, ")")
				}
			}
		}
	}
}
