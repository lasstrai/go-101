package main

func merge(nums1 []int, m int, nums2 []int, n int) {
	j := 0
	for i := 0; i <= m; i++ {
		if nums1[i] > nums2[j] {
			temp := nums2[j]
			nums2[j] = nums1[i]
			nums1[i] = temp
		}
	}
}
