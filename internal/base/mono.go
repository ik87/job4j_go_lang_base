package base

func Mono(nums []int) bool {
	size := len(nums)
	up := true
	down := true

	for i := 0; i < size-1; i++ {
		up = up && nums[i] >= nums[i+1]
		down = down && nums[i] <= nums[i+1]
	}
	return up || down
}
