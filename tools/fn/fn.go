package fn

import "strings"

//截取字符串 start 起点下标 length 需要截取的长度
func Substr(str string, start int, length int) string {
	rs := []rune(str)
	rl := len(rs)
	end := 0

	if start < 0 {
		start = rl - 1 + start
	}
	end = start + length

	if start > end {
		start, end = end, start
	}

	if start < 0 {
		start = 0
	}
	if start > rl {
		start = rl
	}
	if end < 0 {
		end = 0
	}
	if end > rl {
		end = rl
	}

	return string(rs[start:end])
}

func NicknameByEmail(email string) string {
	arr := strings.Split(email, "@")
	pre := strings.Replace(arr[0], Substr(arr[0], len(arr[0]) / 2, len(arr[0])), "***", 1 )
	nk := pre + "@" + arr[1]
	return nk
}
