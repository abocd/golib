package ustrings

import (
	"regexp"
	"strconv"
)

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

	if end < 0 || end > length {
		return ""
	}
	return string(rs[start:end])
}


func SubstrLength(str string,start int, length int)string{
	rs := []rune(str)
	slength := len(rs)
	if start > slength {
		return ""
	}
	if start < 0 {
		start = slength + start
	}

	if length < 0{
		return ""
	}
	end := start + length


	if end > slength {
		end = slength
	}
	return string(rs[start:end])
}


func MatchStringNum(s string)(isNumber bool,num int){
	isNumber,_ = regexp.Match("^\\d+$",[]byte(s))
	if isNumber {
		num,_ = strconv.Atoi(s)
	}
	return
}