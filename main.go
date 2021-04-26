package main

import (
	"LeetCode/slidingwindow"
	"fmt"
)

func main() {
	nums := []int{7,2,4}

	fmt.Println(slidingwindow.MaxSlidingWindow(nums, 2))
}
