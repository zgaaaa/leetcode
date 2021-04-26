package slidingwindow

/*
给你一个整数数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。
你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位。返回滑动窗口中的最大值。

示例 1：
输入：nums = [1,3,-1,-3,5,3,6,7], k = 3
输出：[3,3,5,5,6,7]
解释：
滑动窗口的位置                最大值
---------------               -----
[1  3  -1] -3  5  3  6  7       3
 1 [3  -1  -3] 5  3  6  7       3
 1  3 [-1  -3  5] 3  6  7       5
 1  3  -1 [-3  5  3] 6  7       5
 1  3  -1  -3 [5  3  6] 7       6
 1  3  -1  -3  5 [3  6  7]      7
*/

func MaxSlidingWindow(nums []int, k int) []int {
	res := []int{}
	if nums == nil {
		return res
	}
	win := []int{}
	for i, v := range nums {
		for len(win) > 0 && v > nums[win[len(win)-1]] {
			win = win[:len(win)-1]
		}
		win = append(win, i)
		if win[0] <= i-k {
			win = win[1:]
		}
		if i >= k-1 {
			res = append(res, nums[win[0]])
		}
	}
	return res
}
