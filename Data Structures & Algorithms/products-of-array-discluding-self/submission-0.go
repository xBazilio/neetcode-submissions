func productExceptSelf(nums []int) []int {
	res := make([]int, len(nums))
	pre := 1
	for i := range len(nums) {
		if i == 0 {
			res[i] = pre
		} else {
			pre = pre * nums[i-1]
			res[i] = pre
		}
	}
	post := 1
	for i := len(nums) - 1; i >= 0; i-- {
		if i == len(nums)-1 {
			res[i] = res[i] * post
		} else {
			post = post * nums[i+1]
			res[i] = res[i] * post
		}
	}
	return res
}
