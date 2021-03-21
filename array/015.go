package array

import "sort"

/*
给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？
请你找出所有和为 0 且不重复的三元组。
注意：答案中不可以包含重复的三元组。

示例 1：
输入：nums = [-1,0,1,2,-1,-4]
输出：[[-1,-1,2],[-1,0,1]]

示例 2：
输入：nums = []
输出：[]

示例 3：
输入：nums = [0]
输出：[]
*/

func ThreeSum(nums []int) [][]int {
	// 先排序
	sort.Ints(nums)
	res := [][]int{}
	length := len(nums)
	// 如果长度小于3直接返回空
	if length < 3 {
		return res
	}
	// 对排好序的数组从左往右遍历
	for i, v := range nums {
		// 因为是排好序的,如果固定的值已经大于0,后面的值一定大于0,所以直接返回
		if v > 0 {
			return res
		}
		// 如果当前固定的值与前一个相同,跳过
		if i > 0 && v == nums[i-1] {
			continue
		}
		// 指针初始位置,左指针为固定值的下一个,右指针为最后一个
		l, r := i+1, length-1
		// 移动指针
		for r > l {
			// 求和
			sum := v + nums[l] + nums[r]

			switch {
			case sum == 0:
				// 保存结果
				res = append(res, []int{v, nums[l], nums[r]})
				// 去重,在r > l的前提下,如果下一个值相同,继续挪动
				for r > l && nums[r] == nums[r-1] {
					r--
				}
				for r > l && nums[l] == nums[l+1] {
					l++
				}
				// 移动指针
				l++
				r--
			case sum < 0:
				l++
			case sum > 0:
				r--
			}
		}
	}
	return res
}
