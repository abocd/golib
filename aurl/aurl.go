package aurl

import (
	"net/url"
	"regexp"
)

/**
  补齐相对url为绝对url
  @param pageUrl 页面url
  @param currUrl 需要补齐的url
 */
func CompleteUrl(pageUrl,currUrl string)string{
	u1,_ := url.Parse(currUrl)
	u2,_ := url.Parse(pageUrl)
	u3 := u2.ResolveReference(u1)
	return u3.String()
}

/**
 判断是否为绝对路径
 */
func CheckAbsUrl(currUrl string)bool{
	exp,err2 := regexp.Compile("^(http://|https://|ftp://)")
	if err2 != nil{
		return false
	}
	return exp.Match([]byte(currUrl))
}