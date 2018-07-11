package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	var strs = make(map[string]int)
	for _, str := range strings.Fields(s) {
		strs[str]++
	}
	return strs
}

func main() {
	wc.Test(WordCount)
}
