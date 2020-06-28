// 爬楼梯.
func climbStairs(n int) int {
    // 1.递归.
    // 2.动态规划.
    // 3.动态规划+缓存 时间复杂度O(n)、空间复杂度O(1)
    i := 1
    j := 2
    if (n == 1) {
        return 1
    }
    if (n == 2) {
        return 2
    }
    // res[n] = res[n - 1] + res[n - 2];
    for m := 3; m <= n; m++ {
        i, j = j, i + j
    }
    return j
}