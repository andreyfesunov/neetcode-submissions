func maxArea(heights []int) int {
    left, right, m := 0, len(heights) - 1, 0

    for left < right {
        s := (right - left) * min(heights[left], heights[right])
        m = max(m, s)
    
        if heights[left] < heights[right] {
            left++
        } else {
            right--
        }
    }
    
    return m
}
