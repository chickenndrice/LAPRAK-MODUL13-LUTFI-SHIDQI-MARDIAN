package main

import "fmt"

func main() {
	var n, o int
	fmt.Scan(&n)

	for n > 0 {
		n /= 10
		o = o + 1
		if n <= 0 {
			break
		}
	}
	fmt.Print(o)
}
