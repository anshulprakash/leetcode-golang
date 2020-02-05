package problem0643

import "math"

func findMaxAverage(nums []int, k int) float64 {
	var windowSum, maxSum int
	var windowStart int
	maxSum = math.MinInt32
	for idx := range nums {
		windowSum = windowSum + nums[idx]
		if idx >= k-1 {
			if windowSum > maxSum {
				maxSum = windowSum
			}
			windowSum = windowSum - nums[windowStart]
			windowStart++
		}
	}
	return float64(maxSum) / float64(k)
}
