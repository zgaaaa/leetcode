package main

import (
	"fmt"
)

func main() {
	map1 := map[int]string{1:"q",2:"w"}
	i,ok := map1[1]
	fmt.Println(i,ok)
}