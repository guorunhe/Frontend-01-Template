// 接雨水, 暴力求解.
func trap(height []int) int {
    ans := 0
    for i, v := range height {
        leftMax := 0
        rightMax := 0
        for j := i; j < len(height); j++ {
            if rightMax < height[j] {
                rightMax = height[j]
            }
        }
        for j := i; j >= 0; j-- {
            if leftMax < height[j] {
                leftMax = height[j]
            }
        }

        tmp := leftMax
        if tmp > rightMax {
            tmp = rightMax
        }
        ans += tmp - v
    }
    return ans
}
