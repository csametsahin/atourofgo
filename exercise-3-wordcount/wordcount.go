package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

var m map[string]int

func WordCount(s string) map[string]int {
	m = make(map[string]int)
	stringfields := strings.Fields(s)

	for _, value := range stringfields {
		v, ok := m[value]
		if !ok {
			m[value] = 1
		} else {
			//wtf when i move else up it worked
			m[value] = v + 1
		}

	}

	return m
}

func main() {
	wc.Test(WordCount)
}
