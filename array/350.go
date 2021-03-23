package array

import "sort"

/*
给定两个数组，编写一个函数来计算它们的交集。

示例 1：
输入：nums1 = [1,2,2,1], nums2 = [2,2]
输出：[2,2]

示例 2:
输入：nums1 = [4,9,5], nums2 = [9,4,9,8,4]
输出：[4,9]

说明：
输出结果中每个元素出现的次数，应与元素在两个数组中出现次数的最小值一致。
我们可以不考虑输出结果的顺序。

进阶：
如果给定的数组已经排好序呢？你将如何优化你的算法？
如果 nums1 的大小比 nums2 小很多，哪种方法更优？
如果 nums2 的元素存储在磁盘上，内存是有限的，并且你不能一次加载所有的元素到内存中，你该怎么办？
*/

func Intersect(nums1 []int, nums2 []int) []int {
	// 排序
	sort.Ints(nums1)
	sort.Ints(nums2)
	// 三指针，用于操作nums1，nums2，及结果
	n1, n2, k := 0, 0, 0
	// 循环条件是指针小于数组长度
	for n1 < len(nums1) && n2 < len(nums2) {
		// 因为是排好序的数组,所以小的就向前移
		if nums1[n1] < nums2[n2] {
			n1++
		} else if nums2[n2] < nums1[n1] {
			n2++
		// 相等时就将结果保存
		} else {
			nums1[k] = nums1[n1]
			n1++
			n2++
			k++
		}
	}
	return nums1[:k]
}
