package main

import "fmt"

func main() {
	var target, donasi int
	var lanjut bool
	fmt.Scan(&target, &donasi)

	donatur := 0
	totalDonasi := 0
	for lanjut = true; lanjut;  {
		donatur++
		totalDonasi += donasi
		fmt.Printf("Donatur %d: Menyumbang %d Total terkumpul: %d\n", donatur, donasi, totalDonasi )
		if totalDonasi >= target {
			break
		}
		fmt.Scan(&donasi)
	}

	fmt.Printf("Target tercapai! Total donasi: %d dari %d donatur.", totalDonasi, donatur)
}