package main

import (
	"fmt"
)

/*
	un numero naturale perfetto è uguale alla somma dei suoi divisori propri.
	ex. 6 = 1 + 2 + 3
*/

func main() {
	var n int
	fmt.Print("Inserisci un numero: ")
	fmt.Scan(&n)
	if Perfetto(n) == true {
		fmt.Println(n, " è perfetto")
	} else {
		fmt.Println(n, " non è perfetto")
	}
}

func SommaDivisoriPropri(n int) int {
	var totale int
	var divisori []int
	for i := 1; i < n; i++ {
		if n%i == 0 {
			divisori = append(divisori, i)

		}
	}

	for j := range divisori {
		totale += divisori[j]
	}

	return totale
}

func Perfetto(n int) bool {
	if SommaDivisoriPropri(n) == n {
		return true
	} else {
		return false
	}
}
