// +build ignore

package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z, p := float64(1.0), float64(0.0)
	for math.Abs(p-z) > 1e-6 {
		p = z
		z = z - (z*z-x)/(2*z)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
