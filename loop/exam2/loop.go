package exam2

import "strings"

// 받은 cnt 만큼 str을 붙인 문자열을 반환해주는 함수입니다.
func Repeat(str string, cnt int) string {
	return strings.Repeat(str, cnt)
}
