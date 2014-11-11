// +build ignore

package main

import (
	"fmt"
	"math/cmplx"
)

func Cbrt(x complex128) complex128 {
	z, p := complex128(1), complex128(0)
	for cmplx.Abs(p-z) > 1e-6 {
		p = z
		z = z - (z*z*z-x)/(3*(z*z))
	}
	return z
}

func main() {
	fmt.Println(Cbrt(2))
}
