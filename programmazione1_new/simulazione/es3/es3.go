package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Punto struct {
	etichetta string
	x, y      float64
}

func main() {
	percorso := NuovoTragitto()
	lunghezza := Lunghezza(percorso)
	fmt.Printf("Lunghezza percorso: %.3f\n", lunghezza)

	// Trova il punto oltre la metà
	puntoMedio := PuntoOltreMetà(lunghezza, percorso)
	fmt.Println(puntoMedio)
}

func NuovoTragitto() (tragitto []Punto) {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		// Dividi la riga usando il separatore ';'
		parts := strings.Split(line, ";")

		if len(parts) == 3 {
			var punto Punto
			punto.etichetta = parts[0]

			// Converti x e y da string a float64
			x, errX := strconv.ParseFloat(parts[1], 64)
			y, errY := strconv.ParseFloat(parts[2], 64)

			if errX == nil && errY == nil {
				punto.x = x
				punto.y = y
				tragitto = append(tragitto, punto)
			}
		}
	}

	return tragitto
}

func Distanza(p1, p2 Punto) float64 {
	calcx := p1.x - p2.x
	calcy := p1.y - p2.y
	return math.Sqrt(calcx*calcx + calcy*calcy)
}

func String(p Punto) string {
	return fmt.Sprintf("%s = (%.1f, %.1f)", p.etichetta, p.x, p.y)
}

func Lunghezza(tragitto []Punto) float64 {
	var lunghezzaTotale float64
	for i := 0; i < len(tragitto)-1; i++ {
		lunghezzaTotale += Distanza(tragitto[i], tragitto[i+1])
	}
	return lunghezzaTotale
}

func PuntoOltreMetà(lunghezzaTotale float64, tragitto []Punto) string {
	metà := lunghezzaTotale / 2
	distanzaPercorsa := 0.0

	for i := 0; i < len(tragitto)-1; i++ {
		distanzaTratta := Distanza(tragitto[i], tragitto[i+1])
		distanzaPercorsa += distanzaTratta

		if distanzaPercorsa > metà {
			return fmt.Sprintf("Punto oltre metà: %s", String(tragitto[i+1]))
		}
	}

	// Se non troviamo un punto oltre la metà, restituisci l'ultimo
	return fmt.Sprintf("Punto oltre metà: %s", String(tragitto[len(tragitto)-1]))
}

//distanza euclidea = sqrt(((a.x-b.x)^2+(a.y-b.y)^2))
