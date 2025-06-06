package main

import (
	"fmt"
	"strings"
)

func main() {
	var n int
	fmt.Scan(&n)

	// Se n <= 0, stampa solo il tronco
	if n <= 0 {
		// Tronco 3x3
		for i := 0; i < 3; i++ {
			fmt.Println("+++")
		}
		return
	}

	// Calcola la larghezza massima della chioma
	maxWidth := 2*n - 1

	// Stampa la chioma (triangolo che parte da 1 asterisco in alto)
	for i := 1; i <= n; i++ {
		// Numero di asterischi nella riga i-esima (1, 3, 5, 7, ...)
		stars := 2*i - 1
		// Numero di spazi per centrare rispetto alla larghezza massima
		spaces := (maxWidth - stars) / 2

		// Stampa gli spazi seguiti dagli asterischi
		fmt.Println(strings.Repeat(" ", spaces) + strings.Repeat("*", stars))
	}

	// Stampa il tronco (3x3 centrato rispetto alla chioma)
	trunkSpaces := (maxWidth - 3) / 2
	trunkLine := strings.Repeat(" ", trunkSpaces) + "+++"

	for i := 0; i < 3; i++ {
		fmt.Println(trunkLine)
	}
}
