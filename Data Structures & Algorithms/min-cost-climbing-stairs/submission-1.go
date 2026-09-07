func minCostClimbingStairs(cost []int) int {
	cache := make(map[int]int, len(cost) + 1)

    var dfs func(i int) int
    dfs = func(i int) int {
		v, ok := cache[i]
		if ok {
			return v
		}
        if i >= len(cost) {
            return 0
        }
        cache[i] = cost[i] + min(dfs(i+1), dfs(i+2))
		return cache[i]
    }

    return min(dfs(0), dfs(1))
}
