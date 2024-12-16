package main

import "fmt"

func main() {
	var kata string
	var ulang int
	fmt.Scan(&kata, &ulang)
	hitung := 0
	for selesai := false; !selesai; {
		fmt.Println(kata)
		hitung++
		selesai = (hitung >= ulang)
	}
}