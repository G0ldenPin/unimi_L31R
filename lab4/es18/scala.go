package main

import (
	"fmt"
	"strings"
)

func main() {
	var n int
	fmt.Print("Inserisci il numero di gradini: ")
	fmt.Scan(&n)
	stampaScala(n)

}

func stampaGradino(gradino int) {
	larghezza := 3
	altezza := 2

	for i := 0; i < larghezza-1; i++ {
		fmt.Print("* ")
		fmt.Println()
	}
	for j := 0; j < altezza-1; j++ {
		fmt.Print(strings.Repeat(" ", 2))

	}
}

func stampaScala(gradini int) {
	for i := 0; i < gradini; i++ {
		stampaGradino(i)
	}
}
