package main

import "fmt"

func itoabase(value int, base int) string {
	//for now we can just handle base from 2 to 16 10+26=36 digits and caracter
	if base < 2 || base > 16 {
		return ""
	}
	if value == 0 {
		return "0"
	}
	const digits = "0123456789ABCDEF"
	result := ""

	if value < 0 {
		value = -value
		result = "-"
	}
	for value > 0 {
		conv_reminder := value % base
		result = string(digits[conv_reminder]) + result
		value /= base
	}
	return result
}
func main() {
	fmt.Println(itoabase(255, 10))
	fmt.Println(itoabase(255, 16))
	fmt.Println(itoabase(-255, 16))
}
