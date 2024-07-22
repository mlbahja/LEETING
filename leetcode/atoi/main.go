package main

import (
	//"crypto/internal/boring/sig"
	"fmt"
)

func main() {
	str := "--1223"
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
		result = int(char-'0') + result*10
	}
	fmt.Println(result * sign)
}

/*Write a function that simulates the behaviour of the Atoi function in Go.
Atoi transforms a number represented as a string in a number represented as an int.
Atoi returns 0 if the string is not considered as a valid number.
For this exercise non-valid string chains will be tested.
Some will contain non-digits characters.
For this exercise the handling of the signs + or - does have to be taken into account.
This function will only have to return the int.
 For this exercise the error result of Atoi is not required.*/
