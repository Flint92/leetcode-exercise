from typing import List


def largest_rectangle_in_histogram(heights: List[int]) -> int:
    stk = [-1]
    max_area = 0
    for i in range(0, len(heights) + 1):
        curr_height = 0 if i == len(heights) else heights[i]
        while len(stk) > 1 and curr_height < heights[stk[-1]]:
            h = heights[stk.pop()]
            area = (i - 1 - stk[-1]) * h
            max_area = max(max_area, area)
        stk.append(i)
    return max_area
