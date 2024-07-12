package main

import "fmt"

func isPalindrome(x int) bool {
	var a int
	var copy int
	copy = x
	if copy < 0 {
		return false
	}

	for copy > 0 {
		a = a*10 + copy%10
		copy = copy / 10
	}
	if x != a {
		return false
	} else {
		return true
	}
}

func main() {
	a := isPalindrome(121)
	fmt.Println(a)
}
