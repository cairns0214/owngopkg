package lstrings

// 字符串反转
func ReverseString(s string) string {
	runes := []rune(s)
	for from, to := 0, len(s)-1; from < to; from, to = from+1, to-1 {
		runes[from], runes[to] = runes[to], runes[from]
	}
	return string(runes)
}
