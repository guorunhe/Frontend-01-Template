// 两数和.
func twoSum(nums []int, target int) []int {
    // 解1.
    for i := 0; i < len(nums) - 1; i++ {
        for j := i + 1; j < len(nums); j++ {
            if target == (nums[i] + nums[j]) {
                return []int{i, j}
            }
        }
    }
    return nil
    // 解2.
    m := make(map[int]int)

    for i, v := range nums {
        if k, ok := m[target - v]; ok {
            return []int{k, i}
        }
        m[v] = i
    }

    return nil
}
