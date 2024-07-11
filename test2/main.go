func isPalindrome(x int) bool {
    str := itoa(x)
    result := ""
    if x < 0 {
        return false
    }
    for i := len(str) - 1; i >= 0; i-- {
        result += string(str[i])
    }
    if result == str {
        return true
    }else {
        return false
    }
}
func itoa(x int) string {
    res := ""
    if x < 0 {
        res += "-"
        x = -x
    }
    for x > 0 {
        res += string(rune((x%10)+ '0'))
        x /= 10
    }
    return rev(res)
}
func rev(s string) string {
    re := []rune(s)
    i := 0
    j := len(s) - 1
    if s[i] == '-' {
        i++
    }
    for i < j {
        re[i], re[j] = re[j], re[i]
        i++
        j--
    }
    return string(re)
}