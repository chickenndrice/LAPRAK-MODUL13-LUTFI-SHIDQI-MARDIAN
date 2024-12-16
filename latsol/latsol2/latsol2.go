package main

import (
	"fmt"
	"math"
)

func main() {
	var n, output float64
	fmt.Scanln(&n)
	target := math.Ceil(n) 
	tambah := 0.1                
	output = n          
	for { 
		output += tambah 
		output = math.Round(output*10) / 10 
		fmt.Printf("%.1f\n", output)
		if output >= target { 
			break
		}
	}
}

