package main

import (
	"fmt"
)

func main() {
	str := []string{"ab", "a"}
	fmt.Println(LongestCommonPrefix(str))
}

func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	// 纵向扫描,从前往后遍历所有字符串的每一列
	// 遍历每一列
	for j := 0; j < len(strs[0]); j++ {
		// 从第二个字符串遍历
		for i := 1; i < len(strs); i++ {
			// 边界条件,
			// 1.由于外层循环是按照第0个字符串的长度遍历的,有可能后面的字符串没有第0个字符串长
			// 2.当第i个字符串的第j个字符与第0个字符串的第j个字符不相等的时候	
			// 注意:需要先判断字符串长度再判断字符,否则会越界
			if j == len(strs[i]) || strs[i][j] != strs[0][j] {
				return strs[0][:j]
			}
		}
	}
	return strs[0]
}
