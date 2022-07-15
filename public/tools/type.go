package tools

import (
	"strings"
	"unicode"

	"github.com/mozillazg/go-pinyin"
)

// H is a shortcut for map[string]interface{}
type H map[string]interface{}

// 是否全为中文
func isChinese(s string) bool {
	for _, r := range s {
		if !unicode.Is(unicode.Han, r) {
			return false
		}
	}
	return true
}

// 将中文内容转成拼音
func ConvertToPinYin(s string) (ret string) {
	if isChinese(s) {
		ret = strings.Join(pinyin.LazyConvert(s, nil), "")
	} else {
		ret = s
	}
	return
}
