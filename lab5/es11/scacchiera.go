package main

import (
	"fmt"
)

func main() {
	var dimensione int
	fmt.Print("Inserisci la dimensione: ")
	fmt.Scan(&dimensione)

	StampaScacchiera(dimensione)

}

func StampaRigaInizioAst(lunghezza int) {
	for i := 1; i <= lunghezza; i++ {
		if (i%3 == 0) || (i == 1) {
			fmt.Print("* ")
		} else {
			fmt.Print("+ ")
		}
	}
}

func StampaRigaInizioAdd(lunghezza int) {
	for i := 1; i <= lunghezza; i++ {
		if (i%3 == 0) || (i == 1) {
			fmt.Print("+ ")
		} else {
			fmt.Print("* ")
		}
	}
}

func StampaScacchiera(dimensione int) {
	if dimensione <= 0 {
		return
	} else {
		for i := 2; i <= dimensione+1; i++ {
			if i%2 == 0 {
				StampaRigaInizioAst(dimensione)
				fmt.Println()
			} else {
				StampaRigaInizioAdd(dimensione)
				fmt.Println()
			}
		}
	}
}
