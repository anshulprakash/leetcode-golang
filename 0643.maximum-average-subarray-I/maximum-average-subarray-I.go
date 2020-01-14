package problem0643

func findMaxAverage(nums []int, k int) float64 {
	var sum int
	var res float64
	for idx := 0; idx < k; idx++ {
		sum += nums[idx]
	}
	res = float64(sum) / float64(k)
	for idx := k; idx < len(nums); idx++ {
		sum = sum + nums[idx] - nums[idx-k]
		temp := float64(sum) / float64(k)
		if temp > res {
			res = temp
		}
	}
	return res
}
