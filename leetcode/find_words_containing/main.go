package main

import "fmt"

func main() {
	str := []string{"leet", "code"}
	x := "e"
	res := []int{}
	for i, word := range str {
		for j := 0; j < len(word); j++ {
			if string(word[j]) == x {
				res = append(res, i)
				break
			}
		}
	}
	fmt.Println(res)
}
