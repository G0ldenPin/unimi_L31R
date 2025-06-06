package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	var n int
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)
	numgenerato := r.Intn(100)
	for i := 1; i >= 1; i++ {
		fmt.Scan(&n)
		fmt.Println("Tentativo #", i, ":", n)
		if n < numgenerato {
			fmt.Println("Troppo basso! Riprova!")
		} else if n > numgenerato {
			fmt.Println("Troppo alto! Riprova!")
		} else if n == numgenerato {
			fmt.Println("Hai indovinato in ", i, " tentativi!")
			os.Exit(0)
		}

	}
}
