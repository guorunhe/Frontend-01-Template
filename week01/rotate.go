// 循环数组.
func rotate(nums []int, k int)  {
    // 循环k次，移动  时间O(n * k) 空间O(1)
    // for i := 0; i < k; i++ {
    //     tmp := nums[len(nums) - 1]
    //     for j := len(nums) - 1; j > 0; j-- {
    //         nums[j] = nums[j - 1]
    //     }
    //     nums[0] = tmp
    // }
    count := 0
    for i := 0; count < len(nums); i++ {
        j := i
        pre := nums[j]
        for {
            j = (j + k) % len(nums)
            tmp := nums[j]
            count++
            nums[j] = pre
            pre = tmp
            if j == i {
                break
            }
        }
    }
}
