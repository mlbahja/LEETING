package main

import (
	"os"

	"github.com/01-edu/z01"
)

func printstr(s string) {
	for _, char := range s {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}
func result(s1, s2 string) string {
	var rest []rune
	for _, a := range s1 {
		if containsRune(s2, a) && !containsRune(string(rest), a) {
			rest = append(rest, a)
		}
	}
	return string(rest)
}
func containsRune(s string, r rune) bool {
	for _, char := range s {
		if char == r {
			return true
		}
	}
	return false
}
func main() {
	if len(os.Args) != 3 {
		printstr("Usage non valid")
		return
	}
	str1 := os.Args[1]
	str2 := os.Args[2]
	resultstring := result(str1, str2)
	printstr(resultstring)

}
