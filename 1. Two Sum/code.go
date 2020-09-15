func twoSum(nums []int, target int) []int {
    M := map[int]int{}
    for i, n := range nums {
        if _, exists := M[target - n]; exists {
            return []int{M[target - n], i}
        }
        M[n] = i
    }
    return []int{}
}
