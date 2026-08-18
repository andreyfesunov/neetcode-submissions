func characterReplacement(s string, k int) int {
	res := 0
	set := make(map[byte]struct{})

	for i := 0; i < len(s); i++ {
		set[s[i]] = struct{}{}
	}

	for char := range set {
		count, l := 0, 0
		for r := 0; r < len(s); r++ {
			if s[r] == char {
				count++
			}

			for (r - l + 1) - count > k {
				if s[l] == char {
					count--
				}
				l++
			}

			res = max(res, r - l + 1)
		}
	}

	return res
}
