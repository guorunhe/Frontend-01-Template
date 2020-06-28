// 括号匹配.
func isValid(s string) bool {
    // 暴力法,时间超了.
    // oldS := s
    // for {
    //     s = strings.Replace(s, "[]", "", -1)
    //     s = strings.Replace(s, "()", "", -1)
    //     s = strings.Replace(s, "{}", "", -1)
    //     if s == "" {
    //         return true
    //     }
    //     if s == oldS {
    //         return false
    //     }
    // }
    // return false
    var stack []uint8
    if s == "" {
        return true
    }
    m := map[uint8]uint8 {
        '}' : '{',
        ']' : '[',
        ')' : '(',
    }
    for i := 0; i < len(s); i++ {
        if s[i] == '{' || s[i] == '(' || s[i] == '[' {
            stack = append(stack, s[i])
        } else {
            if len(stack) == 0 {
                return false
            }
            if m[s[i]] != stack[len(stack) - 1] {
                return false
            }

            stack = stack[:len(stack) - 1]
        }
    }

    return len(stack) == 0
}
