func hasDuplicate(nums []int) bool {
    mem := make(map[int]struct{})
	for _, i := range nums {
		if _, ok := mem[i]; ok {
			return true
		}
		mem[i] = struct{}{}
	}
	return false
}
