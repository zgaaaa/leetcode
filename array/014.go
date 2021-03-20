package array

/*
编写一个函数来查找字符串数组中的最长公共前缀。
如果不存在公共前缀，返回空字符串 ""。

示例 1：
输入：strs = ["flower","flow","flight"]
输出："fl"

示例 2：
输入：strs = ["dog","racecar","car"]
输出：""
解释：输入不存在公共前缀。
*/

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
			// 1.当第i个字符串的第j个字符与第0个字符串的第j个字符不相等的时候
			// 2.由于外层循环是按照第0个字符串的长度遍历的,有可能后面的字符串没有第0个字符串长
			if strs[i][j] != strs[0][j] || j > len(strs[i]) {
				return strs[0][:j]
			}
		}
	}
	return strs[0]
}
