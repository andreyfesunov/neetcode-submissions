type entry struct {
	k, v int
}

func topKFrequent(nums []int, k int) []int {
	diff := 1000
	acc := make([]entry, 2001)

	for _, value := range nums {
		key := value + diff
		acc[key] = entry{
			k: acc[key].k + 1, v: value,
		}
	}

	sort.Slice(acc, func(i,j int) bool {
		return acc[i].k < acc[j].k
	})

	result := make([]int, 0, k)
	for _, v := range acc[2001 - k:] {
		result = append(result, v.v)
	}
	return result
}
