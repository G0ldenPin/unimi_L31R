package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Inserisci un testo su più righe (termina con CTRL+D): ")
	a, b := StatisticheParole(LeggiTesto())
	fmt.Println("Statistche ---")
	fmt.Println("Numero parole: ", a, "\n Lunghezza media: ", b)
}

func LeggiTesto() string {
	scanner := bufio.NewScanner(os.Stdin)
	var testo string

	for scanner.Scan() {
		testo = scanner.Text()
	}

	return testo
}

func StatisticheParole(s string) (int, int) {
	parole := strings.Fields(s)
	var media, numero, lunghezza int
	numero = len(parole)
	for i := 0; i < numero; i++ {
		lunghezza += len(parole[i])
		media = lunghezza / numero

	}
	return numero, media
}
