func longestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}

	result := string(s[0])
	for i := 0; i < len(s); i++ {
		// First try odd
		cur := 0
		for j, k := i-1, i+1; j >= 0 && k < len(s); j, k = j-1, k+1 {
			if s[j] != s[k] {
				break
			}
			cur++
		}
		if 2*cur+1 > len(result) {
			result = string(s[i-cur : i+cur+1])
		}

		// Try even, [i, i+1]
		cur = 0
		for j, k := i, i+1; j >= 0 && k < len(s); j, k = j-1, k+1 {
			if s[j] != s[k] {
				break
			}
			cur++
		}
		if 2*cur > len(result) {
			result = string(s[i-cur+1 : i+cur+1])
		}
	}

	return result

}

