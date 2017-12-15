package pager

import (
	"math"
	"net/http"
	"net/url"
	"strconv"
	//"fmt"
)

//import "fmt"

type Pager struct {
	Total     int
	PageSize  int
	Page      int
	TotalPage int
}


func (this *Pager) Offset() int {
	if this.Page < 1 {
		this.Page = 1
	}
	if this.PageSize <= 0 {
		this.PageSize = 1
	}
	return (this.Page - 1) * this.PageSize
}

func (this *Pager) Result(r *http.Request) (pageInfo map[string]string, pageList map[int]string) {
	pageInfo = make(map[string]string)
	pageList = make(map[int]string)
	this.TotalPage = int(math.Ceil(float64(this.Total) / float64(this.PageSize)))
	pageInfo["Total"] = strconv.Itoa(this.Total)
	pageInfo["Page"] = strconv.Itoa(this.Page)
	pageInfo["PageSize"] = strconv.Itoa(this.PageSize)
	pageInfo["TotalPage"] = strconv.Itoa(this.TotalPage)
	pageInfo["FirstPage"] = this.getLink(0, r)
	pageInfo["LastPage"] = this.getLink(this.TotalPage, r)
	if this.Page > 0 {
		pageInfo["PrePage"] = this.getLink(this.Page-1, r)
	} else {
		pageInfo["PrePage"] = ""
	}
	if this.Page < this.TotalPage {
		pageInfo["NextPage"] = this.getLink(this.Page+1, r)
	} else {
		pageInfo["NextPage"] = ""
	}
	for i := this.Page - 4; i <= this.Page+4; i++ {
		if i <= 0 || i > this.TotalPage {
			continue
		}
		pageList[i] = this.getLink(i, r)
	}
	//fmt.Println("|||||",pageInfo)
	return
}

func (this *Pager) getLink(page int, r *http.Request) string {
	urlinfo, _ := url.Parse(r.RequestURI)
	query := urlinfo.Query()
	query.Set("page", strconv.Itoa(page))
	urlinfo.RawQuery = query.Encode()
	return urlinfo.String()
}
