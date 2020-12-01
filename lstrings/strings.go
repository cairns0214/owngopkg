// @Author: lvjinyu
// @Time: 2020/12/1 19:43
package lstrings

// @Author lvjinyu
// @Description 字符串反转
// @Date 20:11 2020/12/1
// @Param s string "需要反转的字符串"
// @return s string "反转后的字符串"
func ReverseString(s string) string {
	runes := []rune(s)
	for from, to := 0, len(s)-1; from < to; from, to = from+1, to-1 {
		runes[from], runes[to] = runes[to], runes[from]
	}
	return string(runes)
}
