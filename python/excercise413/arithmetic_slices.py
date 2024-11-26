from typing import List


def number_of_arithmetic_slices(nums: List[int]) -> int:
    count = 0
    n = len(nums)
    if n < 3:
        return count

    dp = [0] * n
    diff = nums[1] - nums[0]
    for i in range(2, n):
        if nums[i] - nums[i-1] == diff:
            dp[i] = dp[i-1] + 1
            count += dp[i]
        else:
            dp[i] = nums[i] - nums[i-1]
    return  count
