func hash(s []int) string {
	return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(s)), "_"), "[]")
}

func threeSum(nums []int) [][]int {
	c := make(map[int][]int)
	for k, v := range nums {
		c[v] = append(c[v], k)
	}

	result := make([][]int, 0)
	seen := make(map[string]struct{})
	n := len(nums)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			a, b := nums[i], nums[j]
			g, ok := c[-(a + b)]
			if len(g) == 1 && (g[0] == i || g[0] == j) {
				continue
			} else if len(g) == 2 && (g[0] == i && g[1] == j || g[1] == i && g[0] == j) {
				continue
			}

			if ok {
				data := []int{a, b, -(a + b)}
				sort.Ints(data)
				h := hash(data)
				_, ok := seen[h]
				if !ok {
					result = append(result, data)
					seen[h] = struct{}{}
				}
			}
		}
	}

	return result
}
