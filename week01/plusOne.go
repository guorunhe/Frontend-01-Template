// 加一.
func plusOne(digits []int) []int {
    for i := len(digits) - 1; i >= 0; i-- {
        num := digits[i] + 1
        digits[i] = num % 10
        if num < 10 {
            return digits
        }
    }
    res := []int{1}
    return append(res, digits...)
}
