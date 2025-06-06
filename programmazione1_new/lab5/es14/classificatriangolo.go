package main

import (
	"fmt"
	"math"
)

func main() {
	var x1, y1, x2, y2, x3, y3 float64
	const epsilon = 1e-9

	fmt.Print("Inserisci i valori dell'ascissa e dell'ordinata nel punto A: ")
	fmt.Scan(&x1, &y1)
	fmt.Print("Inserisci i valori dell'ascissa e dell'ordinata nel punto B: ")
	fmt.Scan(&x2, &y2)
	fmt.Print("Inserisci i valori dell'ascissa e dell'ordinata nel punto C: ")
	fmt.Scan(&x3, &y3)

	ab := distanza(x1, y1, x2, y2)
	bc := distanza(x2, y2, x3, y3)
	ca := distanza(x3, y3, x1, y1)

	if sonoUguali(ab, bc, epsilon) && sonoUguali(bc, ca, epsilon) && sonoUguali(ca, ab, epsilon) {
		fmt.Println("Il triangolo ABC è equilatero")
		fmt.Println("Lunghezza del lato: ", ab)
	} else if sonoUguali(ab, bc, epsilon) || sonoUguali(bc, ca, epsilon) || sonoUguali(ab, ca, epsilon) {
		fmt.Println("Il triangolo ABC è isoscele")
		if sonoUguali(ab, bc, epsilon) {
			fmt.Println("I lati di lunghezza uguale sono AB e BC")
			fmt.Println("Lunghezza dei lati AB e BC: ", ab)
			fmt.Println("Lunghezza del lato CA: ", ca)

		} else if sonoUguali(bc, ca, epsilon) {
			fmt.Println("I lati di lunghezza uguale sono AB e CA")
			fmt.Println("Lunghezza dei lati BC e CA: ", bc)
			fmt.Println("Lunghezza del lato AB: ", ab)
		} else {
			fmt.Println("I lati di lunghezza uguale sono AB e CA")
			fmt.Println("Lunghezza dei lati AB e CA: ", ab)
			fmt.Println("Lunghezza del lato BC: ", bc)
		}
	} else {
		fmt.Println("Il triangolo ABC è scaleno")
		fmt.Println("Lunghezza del lato AB: ", ab)
		fmt.Println("Lunghezza del lato BC: ", bc)
		fmt.Println("Lunghezza del lato CA: ", ca)
	}
}

func distanza(x1, y1, x2, y2 float64) float64 {
	x := (x1 - x2)
	y := (y1 - y2)
	distanza := math.Sqrt(math.Pow(x, 2) + math.Pow(y, 2))

	return distanza
}

func sonoUguali(a, b, epsilon float64) bool {
	return math.Abs(a-b) < epsilon
}
