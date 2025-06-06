package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	var paroleNuove []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		parole := strings.Fields(input)

		for i := range parole {
			paroleNuove = append(paroleNuove, TrasformaParola(parole[i], i))
		}
	}
	fmt.Println(strings.Join(paroleNuove, " "))
}

func TrasformaParola(parola string, posizione int) (parolaTrasformata string) {
	parolaR := []rune(parola)
	for i, r := range parolaR {
		if posizione%2 == 0 {
			if i%2 == 0 {
				parolaR[i] = unicode.ToUpper(r)
			} else {
				parolaR[i] = unicode.ToLower(r)
			}
		} else {
			if i%2 == 0 {
				parolaR[i] = unicode.ToLower(r)
			} else {
				parolaR[i] = unicode.ToUpper(r)
			}
		}
	}
	return string(parolaR)
}
