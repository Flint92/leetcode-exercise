package arithmetic_slices

func numberOfArithmeticSlices(nums []int) int {
	n := len(nums)
	count := 0

	if n < 3 {
		return count
	}

	dp := make([]int, n)
	diff := nums[1] - nums[0]
	for i := 2; i < n; i++ {
		if nums[i]-nums[i-1] == diff {
			dp[i] = dp[i-1] + 1
			count += dp[i]
		} else {
			dp[i] = nums[i] - nums[i-1]
		}
	}

	return count
}
