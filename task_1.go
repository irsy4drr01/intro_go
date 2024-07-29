package main

import "fmt"

func printSegitiga(n int) {
	for i := 0; i <= n; i++ {
		for j := 0; j < i; j++ {
			fmt.Printf(" ")
		}
		for k := 0; k < (2*(n-i) - 1); k++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
