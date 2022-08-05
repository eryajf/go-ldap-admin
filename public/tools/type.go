package tools

import (
	"regexp"
	"strings"
	"unicode"

	"github.com/mozillazg/go-pinyin"
)

// H is a shortcut for map[string]interface{}
type H map[string]interface{}

// 是否全为英文
func isEnglish(str string) bool {
	match, _ := regexp.MatchString(`^[A-Za-z]+$`, str)
	return match
}

// 是否为英文与数字组合
func isEnglishAndNum(str string) bool {
	match, _ := regexp.MatchString(`^[a-zA-Z0-9_-]+$`, str)
	return match
}

// 是否全为中文
func isChinese(s string) bool {
	for _, r := range s {
		if !unicode.Is(unicode.Han, r) {
			return false
		}
	}
	return true
}

/*
由于名称的各种组合情况都有，在转换成拼音时也遇到各种各样的问题。这里做一下简单说明，以后将不再处理类似兼容问题，目前兼容如下情况。
	1.如果名字中有横杠或者下划线连接，将会删去下划线再处理
	2.全是中文：直接转拼音
	3.全是英文：不进行处理，原文呈现
	4.英文与数字组合，不进行处理，原文呈现
	5.如果是中英混合，那么分以下几种情况
		1.开头是中文，结尾不是中文：进入Convert逻辑第一种
		2.开头不是中文，结尾不是中文：进入Convert逻辑第一种
		3.开头不是中文，结尾是中文：进入Convert逻辑第三种

	如再有其他情况，将不再进行兼容处理！！！
*/

func ConvertToPinYin(src string) string {
	if strings.Contains(src, "-") {
		src = strings.ReplaceAll(src, "-", "")
	}
	if strings.Contains(src, "_") {
		src = strings.ReplaceAll(src, "_", "")
	}
	return Convert(src)
}

// 将中文内容转成拼音
func Convert(src string) string {
	var dst string
	if isChinese(src) { // 全是中文
		return strings.Join(pinyin.LazyConvert(src, nil), "")
	}
	if isEnglish(src) || isEnglishAndNum(src) { // 全是英文,或者为英文与数字组合
		return src
	}

	han := regexp.MustCompile("([\u4e00-\u9fa5]+)")
	srcIndex := han.FindAllStringIndex(src, -1)

	if srcIndex[0][0] == 0 { // 开头是中文
		dst = strings.ReplaceAll(src, src[srcIndex[0][0]:srcIndex[0][1]], strings.Join(pinyin.LazyConvert(src[srcIndex[0][0]:srcIndex[0][1]], nil), "")+"-")
	}
	if srcIndex[0][0] > 0 && srcIndex[0][1] < len(src) { // 中间是中文
		dst = strings.ReplaceAll(src, src[srcIndex[0][0]:srcIndex[0][1]], "-"+strings.Join(pinyin.LazyConvert(src[srcIndex[0][0]:srcIndex[0][1]], nil), "")+"-")
	}

	if srcIndex[0][1] == len(src) { // 结尾是中文
		dst = strings.ReplaceAll(src, src[srcIndex[0][0]:srcIndex[0][1]], "-"+strings.Join(pinyin.LazyConvert(src[srcIndex[0][0]:srcIndex[0][1]], nil), ""))
	}

	dstIndex := han.FindAllStringIndex(dst, -1)
	if len(dstIndex) == 0 {
		return dst
	}

	return Convert(dst)
}
