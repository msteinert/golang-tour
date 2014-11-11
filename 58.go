package main

// +build ignore

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	z, p := float64(1.0), float64(0.0)
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	for math.Abs(p-z) > 1e-6 {
		p = z
		z = z - (z*z-x)/(2*z)
	}
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
