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
	fmt.Println(a, b, c)
	x := (b * b) + (c * c)
	y := a * a

	return x == y
}

func ternepitagoriche(soglia int) {
	for i := soglia; i > 3; i-- {
		if ternapitagoricatrue(i, (i - 1), (i - 2)) {
			fmt.Println("(", i-2, ",", i-1, ",", i, ")")
		} else if (i-2 < 0) || (i-1 < 0) {
			return
		}
	}
}
