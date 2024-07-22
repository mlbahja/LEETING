package main

import "fmt"

func sorting_string1(s1 string) string {
	str1 := []rune(s1)
	for i := 0; i < len(str1); i++ {
		for j := i + 1; j < len(str1); j++ {
			if str1[i] > str1[j] {
				str1[i], str1[j] = str1[j], str1[i]
			}
		}
	}
	return string(str1)
}

func sorting_string2(s2 string) string {
	str2 := []rune(s2)
	for i := 0; i < len(str2); i++ {
		for j := i + 1; j < len(str2); j++ {
			if str2[i] > str2[j] {
				str2[i], str2[j] = str2[j], str2[i]
			}
		}
	}
	return string(str2)
}

func main() {
	s1 := "mustaphay"
	s2 := "msuathap"
	if sorting_string1(s1) == sorting_string2(s2) {
		fmt.Println("true")
		return
	}
	if sorting_string1(s1) != sorting_string2(s2) {
		fmt.Println("false")
		return
	}
}
