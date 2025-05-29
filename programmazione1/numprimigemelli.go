package main

import (
	"fmt"
)

/* due numeri primi p e q sono gemelli se p=q+2 */

func main() {
	var soglia int
	fmt.Print("Inserisci un numero: ")
	fmt.Scan(&soglia)
	if soglia <= 0 {
		fmt.Println("La soglia non è positiva. Presto sarai impiccato.")
	} else {
		nprimigemelli(soglia)
	}

}

func primo(n int) bool {
	if n < 1 {
		return false
	}
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func nprimigemelli(soglia int) {
	var primi []int
	numeri := make([]int, soglia)

	for i := 0; i < soglia; i++ {
		numeri[i] = i + 1
	}

	for _, j := range numeri {
		if primo(j) == true {
			primi = append(primi, j)
		}
	}

	for i := 0; i < len(primi)-1; i++ {
		if primi[i+1]-primi[i] == 2 {
			fmt.Print("(", primi[i], ",", primi[i+1], ") ")
		}
	}

}
