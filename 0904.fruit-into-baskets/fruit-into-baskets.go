package problem0904

import "math"

func totalFruit(tree []int) int {
	var windowStart, windowEnd int
	res := math.MinInt32
	countMap := make(map[int]int)
	for windowEnd < len(tree) {
		rightChar := tree[windowEnd]
		if freq, ok := countMap[rightChar]; ok {
			countMap[rightChar] = freq + 1
		} else {
			countMap[rightChar] = 1
		}
		for len(countMap) > 2 {
			leftChar := tree[windowStart]
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
