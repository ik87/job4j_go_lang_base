package base

func Mono(nums []int) bool {
	size := len(nums)
	dir := 0

	if size == 1 {
		return false
	}
	if nums[0] < nums[size-1] {
		dir = 1
	}
	if nums[0] > nums[size-1] {
		dir = -1
	}

	for i := 1; i < size; i++ {
		if dir == 0 && nums[i-1] == nums[i] {
			continue
		}
		if dir == 1 && nums[i-1] <= nums[i] {
			continue
		}
		if dir == -1 && nums[i-1] >= nums[i] {
			continue
		}
		return false
	}
	return true
}
