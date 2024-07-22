package main

import "fmt"

func ide(s []int, n int) bool {
	for _, num := range s {
		if n == num {
			return false
		}
	}
	return true
}

func main() {
	nums := []int{1, 2, 2, 3, 3, 3, 4, 4, 2, 2}
	k := 4
	s := make([]int, 0)
	for i := 0; i < len(nums)-1; i++ {
		first_element := nums[i]
		next_element := nums[i+1]
		if first_element == next_element && k > 0 && ide(s, first_element) {
			s = append(s, first_element)
			k--
		} else {
			continue
		}
	}
	fmt.Println(s)
	/*
		give an integer array nums and an integer k, return the k most frequent elements within the array
		the best cases are generated that the answer is always unique
		you may return the output in any order
		example: input = [1,2,2,3,3,3], k = 2
		output : [2,3]
	*/
}
