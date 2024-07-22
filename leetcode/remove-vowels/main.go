package main

import (
	"fmt"
)

// removeVowels removes all vowels (a, e, i, o, u) from the given string.
func removeVowels(s string) string {
	vowels := "aeiouAEIOU"
	result := ""

	for _, char := range s {
		isvowl := false
		for _, vowl := range vowels {
			if char == vowl {
				isvowl = true
				break
			}
		}
		if !isvowl {
			result += string(char)
		}
	}
	return result
}

func main() {
	input := "Hello world!"
	output := removeVowels(input)
	fmt.Println(output)
}
