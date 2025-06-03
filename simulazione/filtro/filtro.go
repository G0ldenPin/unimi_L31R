package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		input := scanner.Text()
		elem := strings.Fields(input) // divide per spazi

		for i := 0; i < len(elem); i += 2 {
			letter := elem[i]
			count, _ := strconv.Atoi(elem[i+1])
			fmt.Print(strings.Repeat(letter, count))
		}
	}
}
