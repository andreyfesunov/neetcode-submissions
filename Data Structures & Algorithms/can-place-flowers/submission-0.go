func canPlaceFlowers(flowerbed []int, n int) bool {
    f := make([]int, 0, len(flowerbed)+2)
    f = append(f, 0)
    f = append(f, flowerbed...)
    f = append(f, 0)

    for i := 1; i < len(f)-1; i++ {
        if f[i-1] == 0 && f[i] == 0 && f[i+1] == 0 {
            f[i] = 1
            n--
        }
    }

    return n <= 0
}