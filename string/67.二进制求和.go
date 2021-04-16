package string

/*
给你两个二进制字符串，返回它们的和（用二进制表示）。

输入为 非空 字符串且只包含数字 1 和 0。



示例 1:

输入: a = "11", b = "1"
输出: "100"
示例 2:

输入: a = "1010", b = "1011"
输出: "10101"
*/

func AddBinary(a string, b string) string {
	la, lb := len(a), len(b)
	// 确保a比b长
	if lb > la {
		la, lb = lb, la
		a, b = b, a
	}
	c := []byte(a)
	var temp byte = 0 // 判断是否进一,1为进一
	for i := la - 1; i >= 0; i-- { // 从后往前遍历
		// a[i]是字符编码不是数字,所以要减去字符'0'的编码
		sum := a[i] - '0' + temp 
		if j := lb + i - la; j >= 0 {
			sum += b[j] - '0' 
		}
		temp = sum / 2 // 只有sum为2时,temp才等于1
		c[i] = sum%2 + '0' // c需要存储字符编码,所以要加上'0'
	}
	res := string(c)
	// 特殊处理,当第一位需要进一时
	if temp == 1 {
		res = "1" + res
	}
	return res
}
