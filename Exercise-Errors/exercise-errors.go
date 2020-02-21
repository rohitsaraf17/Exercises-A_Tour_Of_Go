package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0.0 {
		return 0, ErrNegativeSqrt(x)
	}
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
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
