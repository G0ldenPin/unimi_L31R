package main

import (
	"fmt"
)

func main() {
	fmt.Print("Inserisci un numero:")
	var n, scelta int
	fmt.Scan(&n)
	fmt.Println("Scegli che forma stampare: \n 1]quadrato semplice \n 2]quadrato a righe alterne \n 3]quadrato a righe alterne doppie \n 4]quadrato e area interna \n 5]quadrato a colonne alterne \n 6]quadrato a colonne alterne doppie \n 7]quadrato con diagonale")
	fmt.Scan(&scelta)

	if (scelta < 1) || (scelta > 7) {
		fmt.Println("Scelta errata")
	}

	switch scelta {
	case 1:
		quadrato(n)
	case 2:
		quadratorighealterne(n)
	case 3:
		quadratorighealterne2(n)
	case 4:
		quadratoareainterna(n)
	case 5:
		quadratocolonnealterne(n)
	case 6:
		quadratocolonnealterne2(n)
	case 7:
		quadratocondiagonale(n)
	}
}

func quadrato(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}

func quadratorighealterne(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i%2 == 0 {
				fmt.Print("* ")
			} else {
				fmt.Print("+ ")
			}
		}
		fmt.Println()
	}
}

func quadratorighealterne2(n int) {
	for i := 0; i < n; i++ {
		var simbolo string
		switch i % 3 {
		case 0:
			simbolo = "* "
		case 1:
			simbolo = "+ "
		case 2:
			simbolo = "° "
		}
		for j := 0; j < n; j++ {
			fmt.Print(simbolo + " ")
		}
		fmt.Println()
	}
}

func quadratoareainterna(n int) {
	for i := 0; i <= n; i++ {
		for j := 0; j <= n; j++ {
			if (i == 0) || (i == n) || (j == 0) || (j == n) {
				fmt.Print("* ")
			} else {
				fmt.Print("+ ")
			}
		}
		fmt.Println()
	}
}

func quadratocolonnealterne(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if j%2 == 0 {
				fmt.Print("* ")
			} else {
				fmt.Print("+ ")
			}
		}
		fmt.Println()
	}
}

func quadratocolonnealterne2(n int) {
	var count int
	var stampa = "*"
	count = 2
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if count == 2 {
				if stampa == "* " {
					stampa = "+ "
				} else {
					stampa = "* "
				}
				count = 0
			}
			count++
			fmt.Print(stampa)
		}
		stampa = "* "
		count = 0
		fmt.Println()
	}
}

func quadratocondiagonale(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i < j {
				fmt.Print("+ ")
			} else if i == j {
				fmt.Print("° ")
			} else if i > j {
				fmt.Print("* ")
			}
		}
		fmt.Println()
	}
}
