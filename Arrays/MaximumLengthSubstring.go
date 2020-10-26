package Arrays

import "math"

func MaxLengthSubstring(value string, target int) (maxLength int) {
	maxLength = math.MinInt64

	distinctMap := map[rune]int{}
	start := 0

	for end := 0; end < len(value); end++ {
		distinctMap[rune(value[end])] += 1
		for ; len(distinctMap) > target; {
			maxLength = max(maxLength, end-start)
			distinctMap[rune(value[start])]--
			if distinctMap[rune(value[start])] == 0 {
				delete(distinctMap, rune(value[start]))
			}
			start++
		}
	}
	return
}

func max(first, second int) (value int) {
	if first > second {
		value = first
	} else {
		value = second
	}
	return
}
