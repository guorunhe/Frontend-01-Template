// 移动零.
func moveZeroes(nums []int)  []int{
    if nums == nil {
        return nil
    }

    var j int
    for i := 0; i < len(nums); i++ {
        if nums[i] == 0 {
            continue
        }

        nums[i], nums[j] = nums[j], nums[i]
        j++
    }

    return nums
}
