package problem0003

func lengthOfLongestSubstring(s string) int {
	set := make(map[string]int)
	var windowStart, windowEnd, res int
	for windowEnd < len(s) {
		rightChar := s[windowEnd : windowEnd+1]
		if idx, ok := set[rightChar]; ok {
			for windowStart <= idx {
				delete(set, s[windowStart:windowStart+1])
				windowStart++
			}
		}
		set[rightChar] = windowEnd
		if windowEnd-windowStart+1 > res {
			res = windowEnd - windowStart + 1
		}
		windowEnd++
	}

	return res
}
