package binarysearch

/*
给定两个数组，编写一个函数来计算它们的交集。

示例 1：
输入：nums1 = [1,2,2,1], nums2 = [2,2]
输出：[2]

示例 2：
输入：nums1 = [4,9,5], nums2 = [9,4,9,8,4]
输出：[9,4]
*/

func Intersection(nums1 []int, nums2 []int) []int {
	m := map[int]bool{}
	res := []int{}
	for _, v := range nums1 {
		m[v] = true
	}
	for _, v := range nums2 {
		// 错误写法:
		// if _,ok := m[v];ok{ 该写法是判断哈希表m中是否存在v
		if m[v] {
			m[v] = false
			res = append(res, v)
		}
	}
	return res
}
