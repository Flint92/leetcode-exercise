from typing import List


def daily_temperatures(temperatures: List[int]) -> List[int]:
    n = len(temperatures)
    ans = [0] * n

    stk = []
    for i in range(n):
        while len(stk) > 0 and temperatures[stk[-1]] < temperatures[i]:
            ans[stk[-1]] = i - stk[-1]
            stk.pop()
        stk.append(i)

    return ans
