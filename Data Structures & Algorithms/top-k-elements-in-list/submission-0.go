func topKFrequent(nums []int, k int) []int {
	freq := make(map[int]int)
	maxfreq := 0
	for _, n := range nums {
		freq[n]++
		maxfreq = max(maxfreq, freq[n])
	}
	buckets := make([][]int, maxfreq)
	for n, f := range freq {
		buckets[f-1] = append(buckets[f-1], n)
	}
	res := make([]int, 0)
	for i := len(buckets) - 1; i >= 0; i-- {
		bucket := buckets[i]
		for _, n := range bucket {
			res = append(res, n)
			if len(res) == k {
				return res
			}
		}

	}
	return res
}
