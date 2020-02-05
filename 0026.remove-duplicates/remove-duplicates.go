package problem0026

func removeDuplicates(nums []int) int {
	var idx1 int

	for idx2 := 1; idx2 < len(nums); idx2++ {
		if nums[idx1] != nums[idx2] {
			nums[idx1+1] = nums[idx2]
			idx1++
		}
	}

	return idx1 + 1
}
