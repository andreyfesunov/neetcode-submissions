func maxNumberOfBalloons(text string) int {
    countText := make(map[rune]int)
    for _, c := range text {
        countText[c]++
    }

    balloon := map[rune]int{'b': 1, 'a': 1, 'l': 2, 'o': 2, 'n': 1}

    res := len(text)
    for c, need := range balloon {
        if countText[c]/need < res {
            res = countText[c] / need
        }
    }
    return res
}