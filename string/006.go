package string

import "strings"

/*
将一个给定字符串 s 根据给定的行数 numRows ，以从上往下、从左到右进行 Z 字形排列。
比如输入字符串为 "PAYPALISHIRING" 行数为 3 时，排列如下：
P   A   H   N
A P L S I I G
Y   I   R
之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，比如："PAHNAPLSIIGYIR"。
请你实现这个将字符串进行指定行数变换的函数：

string convert(string s, int numRows);

示例 1：
输入：s = "PAYPALISHIRING", numRows = 3
输出："PAHNAPLSIIGYIR"

示例 2：
输入：s = "PAYPALISHIRING", numRows = 4
输出："PINALSIGYAHRPI"
解释：
P     I    N
A   L S  I G
Y A   H R
P     I

示例 3：
输入：s = "A", numRows = 1
输出："A"
*/

func Convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	// 因为go语言中字符串不能改变，所以用切片操作
	b := []byte(s)
	// 结果保存到字符串切片中
	res := make([]string, numRows)
	// 周期
	cyc := numRows*2 - 2
	length := len(b)
	// 遍历字符串
	for i := 0; i < length; i++ {
		// 取余数
		q := i % cyc
		if q < numRows {
			res[q] += string(b[i])
		} else {
			res[cyc-q] += string(b[i])
		}
	}
	// Join函数,合并字符串
	return strings.Join(res, "")
}
