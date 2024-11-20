from typing import List


def dfs(ans: List[List[int]], path: List[int], k: int, n: int, start: int):
    if len(path) >= k or n <= 0:
        if len(path) == k and n == 0:
            ans.append(path[:])
        return

    for i in range(start, 10):
        path.append(i)
        dfs(ans, path, k, n - i, i + 1)
        path.pop()


def combination_sum3(k: int, n: int) -> List[List[int]]:
    ans = []
    dfs(ans, [], k, n, 1)
    return ans
