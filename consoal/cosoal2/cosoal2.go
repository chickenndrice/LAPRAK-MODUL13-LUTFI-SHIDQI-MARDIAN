package main

import "fmt"

func main() {
	var n int
	var lanjut bool
	for lanjut = true; lanjut; {
		fmt.Scan(&n)
		lanjut = n <= 0
	}
	fmt.Printf("%d adalah bilangan bulat positif\n", n)
}