func arrangeCoins(n int) int {
    if n <= 3 {
        if n == 1 {
            return 1
        }
        return n - 1
    }

    l, r := 1, n/2+1
    for l < r {
        mid := (l + r) / 2
        coins := int64(mid) * int64(mid+1) / 2
        if coins <= int64(n) {
            l = mid + 1
        } else {
            r = mid
        }
    }

    return l - 1
}