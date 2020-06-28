// 三数和.
func threeSum(nums []int) [][]int {
    sort.Ints(nums)
    n := len(nums)
    res := make([][]int, 0)
    // ret := make([][]int, 0, 0)
    for i := 0; i < n; i++ {
        if i > 0 && nums[i] == nums[i - 1] {
            continue
        }
        k := n - 1
        for j := i + 1; j < len(nums); j++ {
            if j > i + 1 && nums[j] == nums[j - 1] {
                continue
            }
            for (nums[i] + nums[j] + nums[k]) > 0 && j < k {
                k--
            }
            if j == k {
                break
            }
            if (nums[i] + nums[j] + nums[k] == 0) {
                res = append(res, []int{nums[i], nums[j], nums[k]})
            }
        }
        
    }

    return res
}
