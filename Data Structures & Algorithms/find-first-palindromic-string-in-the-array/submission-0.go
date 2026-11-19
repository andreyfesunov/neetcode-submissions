func firstPalindrome(words []string) string {
    for _, w := range words {
        l, r := 0, len(w)-1
        for w[l] == w[r] {
            if l >= r {
                return w
            }
            l++
            r--
        }
    }
    return ""
}