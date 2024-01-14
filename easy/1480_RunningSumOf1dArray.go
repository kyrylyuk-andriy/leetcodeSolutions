func runningSum(nums []int) []int {
    var sum []int
    var sumInALoop int
    for i := 0; i < len(nums); i ++ {
        sumInALoop += nums[i]
        sum = append(sum, sumInALoop)
    }
    return sum
}
