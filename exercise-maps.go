/*
Exercise: Maps

Implement WordCount. It should return a map of the counts of each ¡°word¡± in the string s. The wc.Test function runs a test suite against the provided function and prints success or failure.

You might find strings.Fields helpful.
*/

package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	stringMap  := make(map[string]int)
	subStrings := strings.Fields(s)

	for _, ssKey := range subStrings {
		ssValue, ssValueExists  := stringMap[ssKey]
		if ssValueExists {
			stringMap[ssKey] = ssValue + 1
		} else {
			stringMap[ssKey] = 1
		}
	} 
	
	return stringMap
}

func main() {
	wc.Test(WordCount)
}
