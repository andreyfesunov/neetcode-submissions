func successfulPairs(spells []int, potions []int, success int64) []int {
    n, m := len(spells), len(potions)
    S := make([]int, n)
    copy(S, spells)
    count := make(map[int]int)
    sort.Ints(spells)
    sort.Ints(potions)

    j := m - 1
    for i := 0; i < n; i++ {
        for j >= 0 && int64(spells[i])*int64(potions[j]) >= success {
            j--
        }
        count[spells[i]] = m - j - 1
    }

    res := make([]int, n)
    for i := 0; i < n; i++ {
        res[i] = count[S[i]]
    }

    return res
}