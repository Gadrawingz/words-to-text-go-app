package main

import (
	"fmt"
	"strings"
)

func main() {
	words := []string{
		"NBA",
		"Basketball",
		"Rap",
		"Arts",
		"Orange",
	}

	// Join each element of the slice
	finalWords := strings.Join(words, ", ")
	fmt.Println(finalWords) // NBA YoungBoy Rapper
}
