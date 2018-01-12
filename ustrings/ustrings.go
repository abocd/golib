package ustrings

import (
	"regexp"
	"strconv"
)

/**
 切割字符串到指定位置
 */
func Substr(str string, start int, end int) string {
	rs := []rune(str)
	length := len(rs)
	if start > length{
		return ""
	}
	if start < 0 {
		start = length + start
	}

	if end < 0{
		end = length + start
	}
	if end < start{
		return ""
	}

	if end < 0 {
		return ""
	}
	if end >= (length - 1){
		end = length - 1
	}
	return string(rs[start:end+1])
}

/**
 切割字符串
 */
func SubstrLength(str string,start int, length int)string{
	return Substr(str,start,start + length - 1)
}

/**
 判断字符串是否是数字
 */
func MatchStringNum(s string)(isNumber bool,num int){
	isNumber,_ = regexp.Match("^\\d+$",[]byte(s))
	if isNumber {
		num,_ = strconv.Atoi(s)
	}
	return
}