package main

func searchInsert(nums []int, target int) int {
	low := 0
	high := len(nums) - 1
	var mid int
	for low <= high {
		mid = low + ((high - low) / 2)
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	if nums[mid] < target {
		return mid + 1
	} else {
		return mid
	}
}
