package main

import (
	"fmt"
)

/*
Due numeri naturali x e y , con x < y , sono detti amichevoli se la somma dei divisori
propri di ciascuno è uguale all’altro; ad esempio (220, 284) è una coppia di amichevoli, essendo
284 = 1 + 2 + 4 + 5 + 10 + 11 + 20 + 22 + 44 + 55 + 110 (dove 1, 2, 4, 5, 10, 11, 20, 22, 44,
55, 110 sono i divisori di 220 ) e 220 = 1 + 2 + 4 + 71 + 142 (dove 1, 2, 4, 71, 142 sono i
divisori di 284 ).
*/

func main() {
	var soglia int
	fmt.Print("Inserisci un numero: ")
	fmt.Scan(&soglia)

	if soglia <= 0 {
		fmt.Println("La soglia non è positiva")
	} else {
		numeriAmichevoli(soglia)
	}
}

func sommaDivisoriPropri(n int) int {
	var totale int
	for i := 1; i < n; i++ {
		if n%i == 0 {
			totale += i
		}
	}
	return totale
}

func sonoAmichevoli(n, m int) bool {
	return sommaDivisoriPropri(n) == m && sommaDivisoriPropri(m) == n && n < m
}

func numeriAmichevoli(limite int) {
	fmt.Print("Numeri amichevoli inferiori a: ", limite)
	for i := 2; i < limite; i++ {
		j := sommaDivisoriPropri(i)
		if j > i && j < limite && sonoAmichevoli(i, j) {
			fmt.Printf("\n (%d, %d)\n", i, j)
		}
	}
}
