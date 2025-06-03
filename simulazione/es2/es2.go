package main

import (
	"bufio"
	"os"
)

type occorrenza struct {
	substring  string
	occorrenza int
}

func main() {
	var substrings []occorrenza
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()

		if len(input)%2 == 0 {
			for i := 0; i < len(input); i++ {
				for j := len(input); i <= 0; j-- {
					if input[i] == input[j] {
						substring := occorrenzeinput[input[i]:input[j]]
						substrings = append(substrings, substring)
					}
				}
			}
		} else {

		}
	}
}
