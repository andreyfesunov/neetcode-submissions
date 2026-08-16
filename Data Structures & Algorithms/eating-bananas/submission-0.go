func minEatingSpeed(piles []int, h int) int {
    l, r := 0, int(1e9)

    for r - l > 1 {
        m := l + (r - l) / 2
        total := 0
        
        for _, v := range piles {
            if m > v {
                total += 1
            } else {
                total += v / m
                if v % m != 0 {
                    total += 1
                }
            }
        }

        if total > h {
            l = m
        } else {
            r = m
        }

        fmt.Println(l, r, m, total)
    }

    return r
}
