// +build ignore

package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	f, f0, f1 := 0, 0, 1
	return func() int {
		f, f0, f1 = f0, f1, f0+f1
		return f
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
