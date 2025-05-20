package main

import (
	"fmt"
)

func main() {
	fmt.Print("Inserisci il voto:")
	var voto int
	fmt.Scan(&voto)
	switch {
	case (voto < 0) || (voto > 100):
		fmt.Println("Errore")
	case voto < 60:
		fmt.Println("Insufficiente")
	case (voto >= 60) && (voto < 70):
		fmt.Println("Sufficiente")
	case (voto >= 70) && (voto < 80):
		fmt.Println("Buono")
	case (voto >= 80) && (voto < 90):
		fmt.Println("Distinto")
	case (voto >= 90) && (voto <= 100):
		fmt.Println("Ottimo")
	}

}
