package daily_temperatures

func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	ans := make([]int, n)

	stk := make([]int, 0)
	for i := 0; i < n; i++ {
		for len(stk) > 0 && temperatures[stk[len(stk)-1]] < temperatures[i] {
			ans[stk[len(stk)-1]] = i - stk[len(stk)-1]
			stk = stk[:len(stk)-1]
		}
		stk = append(stk, i)
	}

	return ans
}
