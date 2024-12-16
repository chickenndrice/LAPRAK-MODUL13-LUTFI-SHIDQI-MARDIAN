package main

import "fmt"

func main() {
	var usn, pw string
	var kond bool
	for kond = false; !kond; {
		fmt.Scan(&usn, &pw)
		kond = usn == "admin" && pw == "admin12345"
	}
	fmt.Println("Selamat anda berhasil login!")
}