func moveZeroes(nums []int) {
	count := 0
	for _, num := range nums {
		if num == 0 {
			count++
		}
	}

	tmp := make([]int, 0, count)
    for _, num := range nums {
        if num != 0 {
            tmp = append(tmp, num)
        }
    }

    for i := 0; i < len(nums); i++ {
        if i < len(tmp) {
            nums[i] = tmp[i]
        } else {
            nums[i] = 0
        }
    }
}