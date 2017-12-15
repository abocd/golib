package aurl

import (
	"bytes"
	"net/url"
	"strings"
	"regexp"
)

/**
  补齐相对url为绝对url
 */
func CompleteUrl(pageUrl,currUrl string)string{
	if len(currUrl) == 0{
		return pageUrl
	}
	if checkAbsUrl(currUrl){
		return currUrl
	}
	var buf bytes.Buffer
	u,err := url.Parse(pageUrl)
	if err != nil{
		return pageUrl
	}
	buf.WriteString(u.Scheme)
	buf.WriteString("://")
	buf.WriteString(u.Host)
	//if u.Port() != "80" && u.Port() != ""{
	//	buf.WriteString(":")
	//	buf.WriteString(u.Port())
	//}
	if err != nil{
		return currUrl
	}
	if currUrl[0] == '/'{
		buf.WriteString(currUrl)
		return buf.String()
	}
	//剩下的相对地址
	index := strings.LastIndex(pageUrl,"/")
	if index == -1{
		return currUrl
	}
	buf.Reset()
	buf.Write([]byte(pageUrl)[:index+1])
	buf.WriteString(currUrl)
	return buf.String()
}

/**
 判断是否为绝对路径
 */
func checkAbsUrl(currUrl string)bool{
	exp,err2 := regexp.Compile("^(http://|https://|ftp://)")
	if err2 != nil{
		return false
	}
	return exp.Match([]byte(currUrl))
}