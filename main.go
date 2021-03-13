package main

import (
	"LeetCode/hashmap"
	"fmt"
)

func main() {
	arr := []int{2, 7, 3, 4, 13, 14}
	a := hashmap.TwoSum(arr, 9)
	fmt.Println(a)
}
