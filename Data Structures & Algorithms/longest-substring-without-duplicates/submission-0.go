func lengthOfLongestSubstring(s string) int {
	m, l := 0, 0
	seen := make(map[rune]int)

	for i, v := range []rune(s) {
		id, ok := seen[v]
		if ok && id >= l {
			l = id + 1
		}
		m = max(m, i - l + 1)
		seen[v] = i
	}

	return m
}
