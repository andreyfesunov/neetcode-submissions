func reverseString(s []byte) {
    tmp := make([]byte, len(s))
    for i := len(s) - 1; i >= 0; i-- {
        tmp[len(s)-1-i] = s[i]
    }
    for i := 0; i < len(s); i++ {
        s[i] = tmp[i]
    }
}