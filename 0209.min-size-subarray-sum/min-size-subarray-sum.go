package problem0209

import "math"

func minSubArrayLen(s int, nums []int) int {
	var res, windowSum, windowStart, windowEnd int
	arrLen := len(nums)
	res = math.MaxInt32
	for windowEnd < arrLen {
		windowSum = windowSum + nums[windowEnd]
		for windowSum >= s && windowStart <= windowEnd {
			if windowEnd-windowStart+1 < res {
				res = windowEnd - windowStart + 1
			}
			windowSum = windowSum - nums[windowStart]
			windowStart++
		}
		windowEnd++
	}
	if res == math.MaxInt32 {
		res = 0
	}
	return res
}
