package main

import (
	"LeetCode/binarysearch"
	"fmt"
)

func main() {
	s := []int{5,7,7,8,8,10}
	n := 8
	b := binarysearch.SearchRange(s,n)
	fmt.Println(b)

}
