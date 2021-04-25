package array2

func RemoveElement(nums []int, val int) int {
	res := 0
	for _, v := range nums {
		if v != val {
			nums[res] = v
			res++
		}
	}
	return res
}