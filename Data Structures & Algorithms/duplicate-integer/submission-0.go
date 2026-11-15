func hasDuplicate(nums []int) bool {
    mem := make(map[int]int)
	for _, i := range nums {
		mem[i]++
		if mem[i] > 1 {
			return true
		}
	}
	return false
}
