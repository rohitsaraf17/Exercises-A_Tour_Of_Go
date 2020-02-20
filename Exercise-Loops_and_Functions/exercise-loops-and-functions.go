package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	z := 1.0
	flag := true
	for flag {
		temp := fmt.Sprintf("%.3f",z)
		z -= (z*z - x) / (2*z)
		z := fmt.Sprintf("%.3f", z)
		if temp == z {
			flag = false
		}
	}
	return z
}

func main() {
	fmt.Println(Sqrt(9))
}
