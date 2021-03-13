package main

import (
	"fmt"
	"leetcode/hashmap"
)

func main() {
	arr := []int{2, 7, 3, 4, 13, 14}
	a := hashmap.twoSum(arr, 9)
	fmt.Println(a)
}
