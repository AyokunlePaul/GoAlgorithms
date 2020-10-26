package Arrays

import "math"

func smallestSubArraySum(target int, nums []int) (result int) {
	result = math.MaxInt64
	currentWindowSum := 0
	start := 0
	for end := 0; end < len(nums); end++ {
		currentWindowSum += nums[end]
		for ; currentWindowSum >= target; {
			result = min(result, end-start+1)
			currentWindowSum -= nums[start]
			start++
		}
	}
	return
}

func min(first, second int) (result int) {
	if first < second {
		result = first
	} else {
		result = second
	}
	return
}
