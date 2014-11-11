// +build ignore

package main

import (
	"code.google.com/p/go-tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	count := make(map[string]int)
	for i := 0; i < len(words); i++ {
		count[words[i]] += 1
	}
	return count
}

func main() {
	wc.Test(WordCount)
}
