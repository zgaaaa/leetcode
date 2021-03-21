package main

import (
	"fmt"
)

func main() {
	slice := []int{2, 2, 3}
	fmt.Println(removeDuplicates(slice))

}

func removeDuplicates(nums []int) int {
	i := 0
	j := 0
	for i < len(nums)-1 {
		for nums[i] == nums[i + 1] {
			nums = append(nums[:i + 1], nums[i + 2:]...)
			j++
		}
	}
	return i + 1
}
