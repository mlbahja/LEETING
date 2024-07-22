package main

import "fmt"

func main() {
	nslice := []int{1, 2, 3, 4, 4, 3, 2, 1}
	n := 3
	res := []int{}
	for i := 0; i <= (n); i++ {
		res = append(res, nslice[i])
		res = append(res, nslice[i+n])
	}
	fmt.Println(res)
}
