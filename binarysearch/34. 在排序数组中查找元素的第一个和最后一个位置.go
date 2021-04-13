package binarysearch

/*
给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。

如果数组中不存在目标值 target，返回 [-1, -1]。

进阶：

你可以设计并实现时间复杂度为 O(log n) 的算法解决此问题吗？


示例 1：

输入：nums = [5,7,7,8,8,10], target = 8
输出：[3,4]
示例 2：

输入：nums = [5,7,7,8,8,10], target = 6
输出：[-1,-1]
示例 3：

输入：nums = [], target = 0
输出：[-1,-1]
*/

func SearchRange(nums []int, target int) []int {
	l, r := -1, -1 // 返回值
	left, right := 0, len(nums)-1
	// 二分法查找开始位置
	for left < right {
		mid := (right + left) / 2
		if nums[mid] < target {
			// [mid+1,right]
			left = mid + 1
		} else if nums[mid] == target {
			// [left,mid],查找mid前面是否还有元素与target相等
			right = mid
		} else {
			// [left,mid-1]
			right = mid - 1
		}
	}
	if left < len(nums) && nums[left] == target {
		l = left
	}
	// 查找结束位置
	left, right = 0, len(nums)-1
	for left < right {
		// 由于是向下取整,避免死循环
		mid := (right + left + 1) / 2
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] == target {
			// [mid,right],查找mid后面是否有元素与target相等
			left = mid
		} else {
			right = mid - 1
		}
	}
	if left < len(nums) && nums[left] == target {
		r = left
	}
	return []int{l, r}
}
