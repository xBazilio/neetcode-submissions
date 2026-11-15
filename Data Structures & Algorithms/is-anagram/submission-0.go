func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	mem := make(map[rune]int)
	for _, r := range s {
		mem[r]++
	}
	for _, r := range t {
		mem[r]--
	}
	for _, i := range mem {
		if i != 0 {
			return false
		}
	}
	return true
}
