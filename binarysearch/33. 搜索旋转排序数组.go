package binarysearch

/*
整数数组 nums 按升序排列，数组中的值互不相同 。
在传递给函数之前，nums在预先未知的某个下标 k（0 <= k < nums.length）上进行了旋转，
使数组变为 [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]]（下标 从 0 开始 计数）。
例如，[0,1,2,4,5,6,7] 在下标 3 处经旋转后可能变为 [4,5,6,7,0,1,2] 。
给你 旋转后的数组 nums 和一个整数 target ，如果 nums 中存在这个目标值 target ，则返回它的下标，否则返回 -1 。

示例 1：
输入：nums = [4,5,6,7,0,1,2], target = 0
输出：4

示例 2：
输入：nums = [4,5,6,7,0,1,2], target = 3
输出：-1

示例 3：
输入：nums = [1], target = 0
输出：-1
*/

func Search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		}
		// 表名mid之前是升序的
		if nums[left] <= nums[mid] {
			// 如果target在left和mid之间
			if nums[left] <= target && target <= nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
			// mid之后是升序的
		} else {
			// target在mid和right之间
			if nums[mid] <= target && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}
