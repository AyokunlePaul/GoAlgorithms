package leetcode

func threeSum(nums []int) [][]int {
	finalArray := make([][]int, 0)

	if len(nums) == 3 {
		if nums[0]+nums[1]+nums[2] == 0 {
			return [][]int{nums}
		}
	}
	for i := 0; i < len(nums)-3; i++ {
		if i == 0 || nums[i] > nums[i+1] {
			start := i + 1
			end := len(nums) - 1

			for {
				sum := nums[i] + nums[start] + nums[end]
				if sum == 0 {
					var tempArray []int
					tempArray = append(tempArray, nums[i])
					tempArray = append(tempArray, nums[start])
					tempArray = append(tempArray, nums[end])
					finalArray = append(finalArray, tempArray)
				}
				if sum > 0 {
					temp := nums[end]
					for {
						end--
						if end == start || nums[end] != temp {
							break
						}
					}
				} else {
					temp := nums[start]
					for {
						start++
						if start == end || nums[start] != temp {
							break
						}
					}
				}
				if start == end {
					break
				}
			}
		}
	}

	return finalArray
}
