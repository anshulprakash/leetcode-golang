package problem0340

import "math"

func lengthOfLongestSubstringKDistinct(s string, k int) int {
	var windowStart, windowEnd int
	res := math.MinInt32
	countMap := make(map[string]int)
	for windowEnd < len(s) {
		rightChar := s[windowEnd : windowEnd+1]
		if freq, ok := countMap[rightChar]; ok {
			countMap[rightChar] = freq + 1
		} else {
			countMap[rightChar] = 1
		}
		for len(countMap) > k {
			leftChar := s[windowStart : windowStart+1]
			if freq := countMap[leftChar]; freq == 1 {
				delete(countMap, leftChar)
			} else {
				countMap[leftChar] = freq - 1
			}
			windowStart++
		}
		if windowEnd-windowStart+1 > res {
			res = windowEnd - windowStart + 1
		}
		windowEnd++
	}
	if res == math.MinInt32 {
		res = 0
	}
	return res
}
