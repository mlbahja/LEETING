package main

import "fmt"

func is_exite(s []int, n int) bool {
	for _, num := range s {
		if num == n {
			return false
		}
	}
	return true
}

func main() {
	nums := []int{1, 2, 2, 3, 3, 3}
	k := 4
	s := make([]int, 0)
	for i := 0; i < len(nums)-1; i++ {
		first_element := nums[i]
		next_element := nums[i+1]
		if first_element == next_element && k > 0 && is_exite(s, first_element) {
			s = append(s, first_element)
			k--
		} else {
			continue
		}

	}
	fmt.Println(s)
}
