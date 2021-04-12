package binarysearch

/*
编写一个高效的算法来判断 m x n 矩阵中，是否存在一个目标值。该矩阵具有如下特性：
每行中的整数从左到右按升序排列。
每行的第一个整数大于前一行的最后一个整数。

示例 1：
输入：matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 3
输出：true

示例 2：
输入：matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 13
输出：false
*/
func SearchMatrix(matrix [][]int, target int) bool {
	l, r := 0, len(matrix)-1
	// 使用二分查找找到target所在的行
	for l <= r {
		mid := (l + r) / 2
		if matrix[mid][0] == target {
			return true
		} else if matrix[mid][0] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	// 避免一种特殊情况
	// 当matrix的长度为1时,r为-1
	if r == -1 {
		r = 0
	}
	left, right := 0, len(matrix[r])-1
	for left <= right {
		mid := (left + right) / 2
		if matrix[r][mid] == target {
			return true
		} else if matrix[r][mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return false
}
