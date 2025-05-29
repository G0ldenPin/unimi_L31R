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
	stampaGradino(n)

}

func stampaGradino(gradino int) {
	larghezza := 3
	altezza := 2
	for x := 1; x <= gradino; x++ {
		for y := 1; y <= larghezza; y++ {
			fmt.Print("*")
			for z := 1; z <= altezza; z++ {
				strings.Repeat(" ", 2)
				if y == z {
					fmt.Println("*")
				}
			}
		}
	}
}

func stampaScala(gradini int) {
}
