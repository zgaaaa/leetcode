package array2

func MoveZeroes(nums []int)  {
	n := 0
	for _, v := range nums {
		if v != 0 {
			nums[n] = v
			n++
		}
	}
	for i := n; i < len(nums); i++ {
		nums[n] = 0
	}
}