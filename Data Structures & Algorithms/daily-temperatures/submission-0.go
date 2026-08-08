func dailyTemperatures(temperatures []int) []int {
	l := len(temperatures)
	result := make([]int, l)
	carry := make([]int, 0)

	if l == 1 {
		return result
	}

	prev := temperatures[0]
	carry = append(carry, 0)

	for index, curr := range temperatures {
		if index == 0 {
			continue
		}

		if prev < curr {
			handled := 0
			for i := len(carry) - 1; i >= 0; i-- {
				cv := carry[i]
				if temperatures[cv] >= curr {
					break
				}

				result[cv] = index - cv

				handled++
			}
			carry = carry[:len(carry) - handled]
		}

		prev = curr
		carry = append(carry, index)
	}

	return result
}
