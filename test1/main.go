package main

import "fmt"

/*func twoSum(nums []int, target int) []int {
}*/

func main() {
	nums := []int{3, 2, 4}
	target := 6
	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i]+nums[j] == target {
				fmt.Println([]int{j, i})
			}
		}
	}
}
