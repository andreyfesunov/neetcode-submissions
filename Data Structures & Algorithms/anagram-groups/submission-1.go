func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string)
	keys := make([]string, 0)

	for _, s := range strs {
		key := sortString(s)
		_, ok := m[key]
		if !ok {
			keys = append(keys, key)
		}
		m[key] = append(m[key], s)
	}

	result := make([][]string, 0, len(keys))

	for _, key := range keys {
		result = append(result, m[key])
	}

	return result
}

func sortString(s string) string {
	r := []rune(s)

	sort.Slice(r, func(i,j int) bool {
		return r[i] < r[j]
	})

	return string(r)
}
