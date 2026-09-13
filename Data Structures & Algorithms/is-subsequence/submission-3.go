func isSubsequence(s string, t string) bool {
	ptr := 0
	for _, v := range []byte(t) {
		if ptr < len(s) && v == s[ptr] {
			ptr++
		}
	}
	return ptr == len(s)
}
