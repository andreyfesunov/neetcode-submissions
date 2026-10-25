func isPerfectSquare(num int) bool {
    for i := 1; i <= num; i++ {
        sq := i * i
        if sq > num {
            return false
        }
        if sq == num {
            return true
        }
    }
    return false
}