func isPalindrome(s string) bool {
	re := regexp.MustCompile(`[^A-Za-z0-9]`)
	s = strings.ToLower(re.ReplaceAllString(s, ""))
	left, right := 0, len(s) - 1
	for right > left {
		if s[right] != s[left] {
			return false
		}
		right--
		left++
	}
	return true
}
