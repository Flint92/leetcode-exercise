package largest_rectangle_in_histogram

func largestRectangleArea(heights []int) int {
	stack := make([]int, 0)
	stack = append(stack, -1) // 第一个柱子的下标是0，默认他前面一个是-1。

	maxArea := 0

	for i := 0; i <= len(heights); i++ {
		var curHeight int
		if i < len(heights) {
			curHeight = heights[i]
		} else {
			curHeight = 0
		}

		for len(stack) > 1 && curHeight < heights[stack[len(stack)-1]] {
			h := heights[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]

			area := (i - 1 - stack[len(stack)-1]) * h

			maxArea = max(maxArea, area)
		}

		stack = append(stack, i)
	}

	return maxArea
}
