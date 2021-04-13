package twopointers

/*
给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。

示例 1：
输入：height = [0,1,0,2,1,0,1,3,2,1,2,1]
输出：6
解释：上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水。

示例 2：
输入：height = [4,2,0,3,2,5]
输出：9
*/

func Trap(height []int) int {
	l, r := 0, len(height)-1
	lmax, rmax, res := 0, 0, 0
	for l < r {
		// 右指针所指的元素大于左指针所指的元素
		if height[l] < height[r] {
			// 左边最高的柱子
			if height[l] >= lmax {
				lmax = height[l]
				// 左指针所指的元素也小于lmax,全部加入水池
			} else {
				res += lmax- height[l]
			}
			l++
		} else {
			if height[r] >= rmax {
				rmax = height[r]
			}else {
				res += rmax- height[r]
			}
			r--
		}
	}
	return res
}
