package aurl

import (
	"testing"
)

func TestCompleteUrl(t *testing.T) {

	if CompleteUrl("http://www.google.com/","1.html") != "http://www.google.com/1.html"{
		t.Error("err1",CompleteUrl("1.html","http://www.google.com/"))
	}
	if CompleteUrl("http://www.google.com/1/","1.html") != "http://www.google.com/1/1.html"{
		t.Error("err2")
	}
	if CompleteUrl("http://www.google.com/1/","/1.html") != "http://www.google.com/1.html"{
		t.Error("err3")
	}
	if CompleteUrl("http://www.google.com/1/2/","../1.html") != "http://www.google.com/1/1.html"{
		t.Error("err4",CompleteUrl("http://www.google.com/1/2/","../1.html"))
	}
}