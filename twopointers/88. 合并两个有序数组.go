package twopointers

/*
给你两个有序整数数组 nums1 和 nums2，请你将 nums2 合并到 nums1 中，使 nums1 成为一个有序数组。
初始化 nums1 和 nums2 的元素数量分别为 m 和 n 。你可以假设 nums1 的空间大小等于 m + n，
这样它就有足够的空间保存来自 nums2 的元素。


示例 1：
输入：nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
输出：[1,2,2,3,5,6]

示例 2：
输入：nums1 = [1], m = 1, nums2 = [], n = 0
输出：[1]
*/

func Merge(nums1 []int, m int, nums2 []int, n int) {
	// 从后向前比较大小,将大的放在nums1的最后
	for p := n + m; n > 0 && m > 0; p-- {
		if nums1[m-1] > nums2[n-1] {
			nums1[p-1] = nums1[m-1]
			m--
		} else {
			nums1[p-1] = nums2[n-1]
			n--
		}
	}
	// 如果nums2中有剩余,说明剩余的几个数是最小的,直接填进nums1
	for ; n > 0; n-- {
		nums1[n-1] = nums2[n-1]
	}
}
