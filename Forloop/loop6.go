package main

import (
	"fmt"
)

func main() {
	n := 10
	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print("kv")
		}
		fmt.Println()
	}
}

// pattern printing
