package main

import (
	"fmt"
)

func main() {
	var s string
	fmt.Scan(&s)
	if palindroma(s) {
		fmt.Println("Palindroma")
	} else {
		fmt.Println("Non palindroma")
	}

}

func palindroma(s string) bool {
	runes := []rune(s)
	length:= len(runes)

	if len(runes)%2 == 0 {
		for i:=0; i<length/2; i++ {
			if runes[i] != runes[length-1-i] {
				return false
			}
		}
	}
	return true
}
