package main

import "fmt"

func main() {
	nums := []int{1, 6, 4, 4, 4, 6, 1, 1, 8, 8}
	result := []int{}
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			result = append(result, nums[i])
		}
		for nums[i] == nums[i+1] {
			i++
			break
		}
	}
	fmt.Println(result)
}
