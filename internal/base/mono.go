package base

func Mono(nums []int) bool {
	size := len(nums)
	add := 0
	n := nums[0]

	if size == 1 {
		return false
	}
	if nums[0] < nums[1] {
		add = 1
	}
	if nums[0] > nums[1] {
		add = -1
	}

	for i := 0; i < size; i++ {
		if n != nums[i] {
			return false
		}
		n += add
	}
	return true
}
