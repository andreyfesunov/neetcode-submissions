func nextGreaterElement(nums1 []int, nums2 []int) []int {
    nums1Idx := make(map[int]int)
    for i, num := range nums1 {
        nums1Idx[num] = i
    }

    res := make([]int, len(nums1))
    for i := range res {
        res[i] = -1
    }

    for i := 0; i < len(nums2); i++ {
        if _, ok := nums1Idx[nums2[i]]; !ok {
            continue
        }
        for j := i + 1; j < len(nums2); j++ {
            if nums2[j] > nums2[i] {
                idx := nums1Idx[nums2[i]]
                res[idx] = nums2[j]
                break
            }
        }
    }

    return res
}