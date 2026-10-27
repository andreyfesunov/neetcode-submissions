func minimumRecolors(blocks string, k int) int {
    countW := 0
    for i := 0; i < k; i++ {
        if blocks[i] == 'W' {
            countW++
        }
    }

    res := countW
    for i := k; i < len(blocks); i++ {
        if blocks[i-k] == 'W' {
            countW--
        }
        if blocks[i] == 'W' {
            countW++
        }
        if countW < res {
            res = countW
        }
    }

    return res
}