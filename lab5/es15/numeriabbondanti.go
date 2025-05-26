package main

import (
	"fmt"
)

func main() {
	var soglia int
	fmt.Print("Inserisci un numero: ")
	fmt.Scan(&soglia)

	if soglia <= 0 {
		fmt.Println("La soglia inserita non è positiva")
	}

	SommaDivisoriPropri(soglia)
	NumeriAbbondanti(soglia)

}

func SommaDivisoriPropri(n int) int {
	var totale int
	//var divisori []int
	for i := 1; i < n; i++ {
		if n%i == 0 {
			totale += i
		}
	}
	return totale
}

func Abbondante(n int) bool {
	if SommaDivisoriPropri(n) > n {
		return true
	} else {
		return false
	}
}

func NumeriAbbondanti(limite int) {
	fmt.Print("Numeri abbondanti: ")
	for i := 1; i < limite; i++ {
		if Abbondante(i) == true {
			fmt.Print(i, " ")
		}
	}
}
