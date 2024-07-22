package main

import "fmt"

/*
	func main() {
		slice := []int{1, 2, 3, 4}
		for i := 0; i < len(slice); i++ {
			for j := i + 1; j < len(slice); j++ {
				if slice[i] == slice[j] {
					fmt.Println("true")
					return
				} else {
					continue
				}
			}
		}
		fmt.Println("false")
	}
*/
func main() {
	s := []int{1, 2, 3, 4}
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			fmt.Println("true")
			return
		}
	}
	fmt.Println("false")
}
