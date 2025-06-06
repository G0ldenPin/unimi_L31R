package main

import (
	"fmt"
)

func main() {
	var s string
	fmt.Scan(&s)
	invertita := invertiStringa(s)
	fmt.Println(invertita)
}

func invertiStringa(s string) string {
	/*var invs []rune
	for i := len(s) - 1; i >= 0; i-- {
		invs = append(invs, rune(s[i]))
	}

	return string(invs)*/
	if len(s) == 0 {
		return ""
	}
	return invertiStringa(s[1:]) + string(s[0])
}

/*Perché serve if s == "" { return "" } nella funzione ricorsiva?

In una funzione ricorsiva come:
func InvertiStringa(s string) string {
	if s == "" {
		return ""
	}
	return InvertiStringa(s[1:]) + string(s[0])
}

questa riga:
if s == "" {
	return ""
}

è fondamentale perché quando s è vuota (""), non puoi più accedere a s[0] o s[1:] → causerebbe un errore di "slice bounds out of range".
Serve a fermare la ricorsione e cominciare a risalire nella catena delle chiamate.
In breve l'if s == "" { return "" }:

Controlla se s è una stringa vuota.
Serve per terminare la ricorsione in modo sicuro.
Restituisce una stringa vuota come punto di partenza per ricostruire la stringa
invertita.*/
