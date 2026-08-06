func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sb, tb := []byte(s), []byte(t)
	sort.Slice(sb, func(i, j int) bool {
		return sb[i] < sb[j]
	})
	sort.Slice(tb, func(i, j int) bool {
		return tb[i] < tb[j]
	})
	
	for i := range sb {
		if sb[i] != tb[i] {
			return false
		}
	}

	return true
}
