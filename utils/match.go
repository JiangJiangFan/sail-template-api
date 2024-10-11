package utils

import (
	"github.com/dlclark/regexp2"
)

func MatchText(text string, reg string, key string) (bool, string) {
	re := regexp2.MustCompile(reg, 0)

	// 查找匹配项
	match, _ := re.FindStringMatch(text)
	if match != nil {
		capture := match.GroupByName(key)
		return true, capture.String()
	}
	return false, "未匹配"
}
