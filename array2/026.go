package array2

func RemoveDuplicates1(nums []int) int {
	res := 1
	for i, v := range nums {
		if i > 0 && v != nums[i-1] {
			nums[res] = v
			res++
		}
	}
	return res
}