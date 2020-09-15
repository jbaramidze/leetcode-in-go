func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func lengthOfLongestSubstring(s string) int {
	maxSofar := 0
	start := 0
	m := map[string]int{}

	for i, c := range s {
		idx, ok := m[string(c)]
		if ok {
			start = Max(start, idx+1)
		}
		maxSofar = Max(maxSofar, i-start+1)
		m[string(c)] = i
	}

	return maxSofar
}

