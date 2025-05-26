package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int
	fmt.Println("Inserisci due numeri interi: ")
	fmt.Scan(&x, &y)

	fmt.Println("Maggiore: ", maggiore(x, y))
	fmt.Println("Minore: ", minore(x, y))
	fmt.Println("Somma: ", somma(x, y))
	fmt.Println("Differenza: ", differenza(x, y))
	fmt.Println("Prodotto: ", prodotto(x, y))
	fmt.Println("Divsione: ", divisione(x, y))
	fmt.Println("Valore medio: ", valoremedio(x, y))
	fmt.Println("Potenza (for): ", potenzafor(x, y))
	fmt.Println("Potenza (math): ", potenzamath(x, y))
}

func maggiore(x, y int) int {
	var maggiore int
	if x > y {
		maggiore = x
	} else {
		maggiore = y
	}
	return maggiore
}

func minore(x, y int) int {
	var minore int
	if x < y {
		minore = x
	} else {
		minore = y
	}
	return minore
}

func somma(x, y int) int {
	somma := x + y
	return somma
}

func differenza(x, y int) int {
	differenza := maggiore(x, y) - minore(x, y)
	return differenza
}

func prodotto(x, y int) int {
	prodotto := x * y
	return prodotto
}

func divisione(x, y int) float64 {
	divisione := float64(x) / float64(y)
	return divisione
}

func valoremedio(x, y int) float64 {
	media := float64((x + y) / 2)
	return media
}

func potenzafor(x, y int) int {
	var potenzaf int
	potenzaf = x
	for i := 0; i < y-1; i++ {
		potenzaf *= x
	}
	return potenzaf
}

func potenzamath(x, y int) float64 {
	potenzam := math.Pow(float64(x), float64(y))
	return potenzam
}
