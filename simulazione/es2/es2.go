package main

import (
	"fmt"
	"os"
	"sort"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run esercizio_2.go <stringa>")
		return
	}

	s := os.Args[1]
	substrings := make(map[string]int)

	n := len(s)
	for i := 0; i < n-2; i++ {
		for j := i + 2; j < n; j++ {
			if s[i] == s[j] {
				sub := s[i : j+1]
				substrings[sub]++
			}
		}
	}

	if len(substrings) == 0 {
		return // non stampare nulla
	}

	// Raccogli tutte le sottostringhe in una slice
	type entry struct {
		value string
		count int
	}
	var results []entry
	for k, v := range substrings {
		results = append(results, entry{value: k, count: v})
	}

	// Ordina per lunghezza decrescente
	sort.Slice(results, func(i, j int) bool {
		if len(results[i].value) == len(results[j].value) {
			return results[i].value < results[j].value // ordine alfabetico se stessa lunghezza
		}
		return len(results[i].value) > len(results[j].value)
	})

	// Stampa i risultati
	for _, e := range results {
		fmt.Printf("%s -> Occorrenze: %d\n", e.value, e.count)
	}
}
