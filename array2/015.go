package array2

import "sort"

// 给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？
// 请你找出所有和为 0 且不重复的三元组。
// 注意：答案中不可以包含重复的三元组。
// 示例 1：
// 输入：nums = [-1,0,1,2,-1,-4]
// 输出：[[-1,-1,2],[-1,0,1]]

func ThreeSum(nums []int) [][]int {
	sort.Ints(nums)
	res := [][]int{}
	length := len(nums)
	if length < 3 {
		return res
	}
	for i, v := range nums {
		if v > 0 {
			return res
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l, r := i+1, length-1
		for l < r {
			sum := v + nums[l] + nums[r]
			switch {
			case sum == 0:
				res = append(res, []int{v, nums[l], nums[r]})
				for r > l && nums[r] == nums[r-1] {
					r--
				}
				for r > l && nums[l] == nums[l+1] {
					l++
				}
				r--
				l++
			case sum < 0:
				l++
			case sum > 0:
				r--
			}

		}
	}
	return res
}
