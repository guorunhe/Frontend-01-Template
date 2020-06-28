// 盛水最多容器.
func maxArea(height []int) int {
    // 找出面积最大O(n^2)
    res := 0
    i := 0
    j := len(height) - 1

    for i != j {
        h1, h2 := height[i], height[j]
        size := (j - i) * min(h1, h2)
        res = max(size, res)
        if (h1 < h2) {
            i++
        } else {
            j--
        }
    }

    return res
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func min(a, b int) int {
    if a > b {
        return b
    }
    return a
}
