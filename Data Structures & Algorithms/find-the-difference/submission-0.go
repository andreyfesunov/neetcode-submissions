func findTheDifference(s string, t string) string {
    countS := make([]int, 26)
    countT := make([]int, 26)

    for _, c := range s {
        countS[c-'a']++
    }
    for _, c := range t {
        countT[c-'a']++
    }

    for i := 0; i < 26; i++ {
        if countT[i] > countS[i] {
            return string(rune('a' + i))
        }
    }
    return " "
}