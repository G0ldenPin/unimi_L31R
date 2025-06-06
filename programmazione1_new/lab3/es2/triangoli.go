package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		for x := 0; x < i; x++ {
			fmt.Print(" ")
		}
		for j := n; j > i; j-- {
			if (j == n) || (j == (i + 1)) || (i == 0) {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}

		if i != (n - 1) {
			fmt.Println()
		}

	}

	for k := 0; k < n; k++ {
		for i := 0; i < (n - 1); i++ {
			fmt.Print(" ")
		}

		for l := 0; l < k; l++ {
			if (l == 0) || (l == (k - 1)) || (k == (n - 1)) {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
