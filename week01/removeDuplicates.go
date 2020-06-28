// 删除数组重复项.
func removeDuplicates(nums []int) int {
    current := 0
    for i := 0; i < len(nums) - 1; i++ {
        if nums[i] == nums[i + 1] {
            continue
        } else {
            current++
            nums[current] = nums[i + 1]
        }
    }
    return current + 1;
}
