func convert(s string, numRows int) string {
    if numRows == 1 {
        return s
    }
	M := make([]string, numRows)

	c := 0
	inc := true
	for _, p := range s {
		M[c] += string(p)
		if inc == true {
			if c == numRows-1 {
				c--
				inc = false
			} else {
				c++
			}
		} else {
			if c == 0 {
				c++
				inc = true
			} else {
				c--
			}
		}
	}

	res := strings.Join(M, "")
	return res

}
