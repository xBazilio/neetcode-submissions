func twoSum(nums []int, target int) []int {
    mem := make(map[int]int)
	var v2 int
	for k, v := range nums {
		v2 = target - nums[k]

		if k2, ok := mem[v2]; ok {
			return []int{k2, k}
		}
		mem[v] = k
	}
	return nil
}
