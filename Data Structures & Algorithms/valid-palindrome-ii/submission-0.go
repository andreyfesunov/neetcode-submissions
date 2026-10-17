func validPalindrome(s string) bool {
    isPalindrome := func(str string) bool {
        left, right := 0, len(str)-1
        for left < right {
            if str[left] != str[right] {
                return false
            }
            left++
            right--
        }
        return true
    }

    l, r := 0, len(s)-1
    for l < r {
        if s[l] != s[r] {
            skipL := s[l+1 : r+1]
            skipR := s[l:r]
            return isPalindrome(skipL) || isPalindrome(skipR)
        }
        l++
        r--
    }

    return true
}