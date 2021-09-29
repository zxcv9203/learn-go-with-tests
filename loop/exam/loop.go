package exam

// 받은 cnt 만큼 str을 붙인 문자열을 반환해주는 함수입니다.
func Repeat(str string, cnt int) string {
	var repeated string
	for i := 0; i < cnt; i++ {
		repeated += str
	}
	return repeated
}
