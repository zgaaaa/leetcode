package array2

import "sort"

func Intersect(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	n1, n2, k := 0, 0, 0
	for n1 < len(nums1) && n2 < len(nums2) {
		if nums1[n1] < nums2[n2] {
			n1++
		} else if nums1[n1] > nums2[n2] {
			n2++
		} else {
			nums1[k] = nums1[n1]
			n1++
			n2++
			k++
		}
	}
	return nums1[:k]
}
