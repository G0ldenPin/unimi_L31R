package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	LeggiEStampa(scanner)
}

func LeggiEStampa(scanner *bufio.Scanner) {
	for scanner.Scan() {
		righe := scanner.Text()
		fmt.Println(righe)
		LeggiEStampa(scanner)
	}
}
