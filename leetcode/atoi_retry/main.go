package main

import "fmt"

func main() {
	str := "--123456789"
	result := 0
	sign := 1
	for i, char := range str {
		if i == 0 && (char == '-' || char == '+') {
			if char == '-' {
				sign = -sign
			}
			continue
		}
		if char < '0' || char > '9' {
			fmt.Println(0)
			return
		}
		result = result*10 + int(char-'0')
	}
	fmt.Println(result * sign)
}
/*
Write a function that takes two arguments:

s: a numeric string in a given base.
base: a string representing all the different digits that can represent a numeric value.
And return the integer value of s in the given base.

If the base is not valid it returns 0.

Validity rules for a base :

A base must contain at least 2 characters.
Each character of a base must be unique.
A base should not contain + or - characters.
String number must contain only elements that are in base.

Only valid string numbers will be tested.

The function does not have to manage negative numbers.

Expected function
func AtoiBase(s string, base string) int {

}
Usage
Here is a possible program to test your function :

package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.AtoiBase("125", "0123456789"))
	fmt.Println(piscine.AtoiBase("1111101", "01"))
	fmt.Println(piscine.AtoiBase("7D", "0123456789ABCDEF"))
	fmt.Println(piscine.AtoiBase("uoi", "choumi"))
	fmt.Println(piscine.AtoiBase("bbbbbab", "-ab"))
}
*/