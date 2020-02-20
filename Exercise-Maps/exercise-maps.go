package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	arr := strings.Fields(s)
	i := len(arr)
	m := make(map[string]int)
	for i > 0 {
		i--
		m[arr[i]] = m[arr[i]] + 1
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
