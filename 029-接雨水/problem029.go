package main

/*
方法三：双指针
动态规划的做法中，需要维护两个数组 \textit{leftMax}leftMax 和 \textit{rightMax}rightMax，因此空间复杂度是 O(n)O(n)。是否可以将空间复杂度降到 O(1)O(1)？

注意到下标 ii 处能接的雨水量由 \textit{leftMax}[i]leftMax[i] 和 \textit{rightMax}[i]rightMax[i] 中的最小值决定。由于数组 \textit{leftMax}leftMax 是从左往右计算，数组 \textit{rightMax}rightMax 是从右往左计算，因此可以使用双指针和两个变量代替两个数组。

维护两个指针 \textit{left}left 和 \textit{right}right，以及两个变量 \textit{leftMax}leftMax 和 \textit{rightMax}rightMax，初始时 \textit{left}=0,\textit{right}=n-1,\textit{leftMax}=0,\textit{rightMax}=0left=0,right=n−1,leftMax=0,rightMax=0。指针 \textit{left}left 只会向右移动，指针 \textit{right}right 只会向左移动，在移动指针的过程中维护两个变量 \textit{leftMax}leftMax 和 \textit{rightMax}rightMax 的值。

当两个指针没有相遇时，进行如下操作：

使用 \textit{height}[\textit{left}]height[left] 和 \textit{height}[\textit{right}]height[right] 的值更新 \textit{leftMax}leftMax 和 \textit{rightMax}rightMax 的值；

如果 \textit{height}[\textit{left}]<\textit{height}[\textit{right}]height[left]<height[right]，则必有 \textit{leftMax}<\textit{rightMax}leftMax<rightMax，下标 \textit{left}left 处能接的雨水量等于 \textit{leftMax}-\textit{height}[\textit{left}]leftMax−height[left]，将下标 \textit{left}left 处能接的雨水量加到能接的雨水总量，然后将 \textit{left}left 加 11（即向右移动一位）；

如果 \textit{height}[\textit{left}] \ge \textit{height}[\textit{right}]height[left]≥height[right]，则必有 \textit{leftMax} \ge \textit{rightMax}leftMax≥rightMax，下标 \textit{right}right 处能接的雨水量等于 \textit{rightMax}-\textit{height}[\textit{right}]rightMax−height[right]，将下标 \textit{right}right 处能接的雨水量加到能接的雨水总量，然后将 \textit{right}right 减 11（即向左移动一位）。

当两个指针相遇时，即可得到能接的雨水总量。
 */

func trap(height []int) (ans int) {
	left, right := 0, len(height)-1
	leftMax, rightMax := 0, 0
	for left < right {
		leftMax = max(leftMax, height[left])
		rightMax = max(rightMax, height[right])
		if height[left] < height[right] {
			ans += leftMax - height[left]
			left++
		} else {
			ans += rightMax - height[right]
			right--
		}
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
