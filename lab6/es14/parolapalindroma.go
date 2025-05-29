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
	var check bool
	if len(s)%2 == 0 {
		for i := 0; i < len(s)/2; i++ {
			for j := len(s) / 2; j >= 0; j-- {
				if s[i] == s[j] {
					check = true
				}

			}
		}
	} else {

	}
	return check
}
