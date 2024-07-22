package main

import "fmt"

func main() {
	nums := []int{1, 2, 2, 3, 3, 3}
	res := []int{}
	for i := 0; i < len(nums); i++ {
		for j := 1; j < len(nums)-1; j++ {
			if nums[i] == nums[j] {
				res = append(res, nums[i])
				// res = append(res, nums[i])
				// continue
				//break
			}
		}
	}
	fmt.Println(res)
}
