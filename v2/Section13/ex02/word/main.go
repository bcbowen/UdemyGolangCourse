package word

import "strings"

// no need to write an example for this one
// writing a test for this one is a bonus challenge; harder

// UseCount returns totals for each word in a given string
func UseCount(s string) map[string]int {
	xs := strings.Fields(s)
	m := make(map[string]int)
	for _, v := range xs {
		m[v]++
	}
	return m
}

// Count returns the number of words in a given string
func Count(s string) int {
	//return len(strings.Split(s, " "))
	return len(strings.Fields(s))
}
