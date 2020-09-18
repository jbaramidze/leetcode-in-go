func reverse(x int) int {
    if x == -1 << 31 {
        return 0
    }
    if x == 0 {
        return 0
    }
    
    neg := false
    if x < 0 {
        x = -1 * x
        neg = true
    }
    
    result := 0
    for x > 0 {
        if (result > ((1 << 31) - (x%10))/10) {
            return 0
        }
        result = result * 10 + (x%10)
        x = x / 10
    }

    
    if neg == true {
        return -1 * result
    }
    
    return result
}
