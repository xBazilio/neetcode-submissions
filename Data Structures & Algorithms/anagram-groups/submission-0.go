func groupAnagrams(strs []string) [][]string {
	mapped := make(map[[26]int][]string)
	for _, str := range strs {
		var count [26]int

		for _, s := range str {
			count[int(s) - int('a')] += 1
		}

		mapped[count] = append(mapped[count], str)
	}

	res := make([][]string, 0, len(mapped))
	for _, vals := range mapped {
		res = append(res, vals)
	}

	return res
}
