package search_engine

import "strings"

// Deprecated / Not In Use
func NumTokens(tokens []string) map[string]int {
	dictionary := make(map[string]int)
	for _, token := range tokens {
		dictionary[token] = dictionary[token] + 1
	}
	return dictionary
}

func Tokenize(input string) []string {
	return strings.Fields(strings.ToLower(input))
}
