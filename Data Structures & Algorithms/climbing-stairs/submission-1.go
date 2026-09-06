func climbStairs(n int) int {
    cache := make(map[int]int, n + 1)
	
	var dfs func(i int) int
	dfs = func(i int) int {
		v, ok := cache[i]
		if ok {
			return v
		}
		if i == n {
			return 1
		}
		if i > n {
			return 0
		}
		cache[i] = dfs(i + 1) + dfs(i + 2)
		return cache[i]
	}

	return dfs(0)
}
