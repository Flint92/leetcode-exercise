package combination_sum_iii

func combinationSum3(k, n int) [][]int {
	ans := make([][]int, 0)

	var dfs func([]int, int, int) []int
	dfs = func(path []int, num int, start int) []int {
		if len(path) >= k || num <= 0 {
			// 找到一组合适的
			if len(path) == k && num == 0 {
				ans = append(ans, append([]int(nil), path...))
			}
			return path
		}

		for i := start; i <= 9; i++ {
			path = append(path, i)
			path = dfs(path, num-i, i+1)
			path = path[:len(path)-1]
		}

		return path
	}

	dfs([]int{}, n, 1)
	return ans
}
