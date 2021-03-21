package array

/*
给你一个有序数组 nums ，请你 原地 删除重复出现的元素，使每个元素 只出现一次 ，返回删除后数组的新长度。
不要使用额外的数组空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。

示例 1：
输入：nums = [1,1,2]
输出：2, nums = [1,2]
解释：函数应该返回新的长度 2 ，并且原数组 nums 的前两个元素被修改为 1, 2 。不需要考虑数组中超出新长度后面的元素。

示例 2：
输入：nums = [0,0,1,1,1,2,2,3,3,4]
输出：5, nums = [0,1,2,3,4]
解释：函数应该返回新的长度 5 ， 并且原数组 nums 的前五个元素被修改为 0, 1, 2, 3, 4 。不需要考虑数组中超出新长度后面的元素。
*/
// 解法一:由于是用append函数进行删除,所以效率低
func RemoveDuplicates1(nums []int) int {
	i := 0
	for i < len(nums) {
		if nums[i] == nums[i+1] {
			nums = append(nums[:i+1], nums[i+2:]...)
		} else {
			i++
		}
	}
	return i
}

// 解法二:不删除只交换位置
func RemoveDuplicates2(nums []int) int {
	length := len(nums)
	res := 1
	// 从第二个开始遍历
	for i := 1; i < length; i++ {
		// 如果当前值与前一个相同
		if nums[i] != nums[i - 1] {
			// 
			nums[res] = nums[i] 
			res++
		}
	}
	return res
}
