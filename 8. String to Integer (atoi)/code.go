func myAtoi(str string) int {
	i := 0
	for i < len(str) && str[i] == ' ' {
		i++
	}

	minus := false
	if i < len(str) && str[i] == '-' {
		minus = true
		i++
	} else if i < len(str) && str[i] == '+' {
		i++
	}

	r := 0
	for i < len(str) && str[i] == '0' {
		i++
	}

	// fist digit manually
	if i < len(str) && str[i] >= '0' && str[i] <= '9' {
		r += int(str[i] - '0')
		i++
	}

	if minus == true {
		r *= -1
	}

	for i < len(str) && str[i] >= '0' && str[i] <= '9' {
		if r > 0 {
			if r > math.MaxInt32/10 {
				return math.MaxInt32
			}
			r *= 10
			if r > math.MaxInt32-int(str[i]-'0') {
				return math.MaxInt32
			}
			r += int(str[i] - '0')
		} else {
			if r < math.MinInt32/10 {
				return math.MinInt32
			}
			r *= 10
			if r < math.MinInt32+int(str[i]-'0') {
				return math.MinInt32
			}
			r -= int(str[i] - '0')
		}

		i++
	}

	return r
}

