package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// rimuoviCifreGreedy implementa l'algoritmo greedy per rimuovere cifre
func rimuoviCifreGreedy(n string, d int) string {
	var stack []byte
	var rimozioniFatte int

	// Processa ogni cifra
	for i := 0; i < len(n); i++ {
		cifra := n[i]

		// Rimuovi cifre più grandi dalla fine dello stack
		for len(stack) > 0 && rimozioniFatte < d && stack[len(stack)-1] > cifra {
			stack = stack[:len(stack)-1] // pop
			rimozioniFatte++
		}

		stack = append(stack, cifra) // push
	}

	// Se non abbiamo fatto abbastanza rimozioni, rimuovi dalla fine
	for rimozioniFatte < d {
		if len(stack) > 0 {
			stack = stack[:len(stack)-1]
			rimozioniFatte++
		}
	}

	return string(stack)
}

// trovaNumeroMinimo trova il più piccolo numero ottenibile rimuovendo d cifre da n
func trovaNumeroMinimo(n string, d int) int64 {
	// Se dobbiamo rimuovere tutte o più cifre di quelle disponibili
	if d >= len(n) {
		return 0
	}
	// Applica l'algoritmo greedy
	risultatoStr := rimuoviCifreGreedy(n, d)
	risultatoStr = strings.TrimLeft(risultatoStr, "0")
	// Converte il risultato in intero
	numeroFinale, _ := strconv.ParseInt(risultatoStr, 10, 64)
	return numeroFinale
}

func main() {
	n, d := os.Args[1], os.Args[2]
	dInt, _ := strconv.Atoi(d)
	risultato := trovaNumeroMinimo(n, dInt)
	fmt.Printf("numero migliore: %d\n", risultato)
}
