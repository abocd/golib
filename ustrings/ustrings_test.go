package ustrings

import (
	"testing"
)

func TestMatchStringNum(t *testing.T) {
	isNumber,num := MatchStringNum("123")
	if !isNumber{
		t.Error(num)
	}
	isNumber2,num2 := MatchStringNum("1a23")
	if isNumber2{
		t.Error(num2,isNumber2)
	}
	isNumber3,num3 := MatchStringNum("aaaa")
	if isNumber3{
		t.Error(num3,isNumber3)
	}
}


func TestSubstr(t *testing.T) {
	var s = "this is a chinese people"
	if Substr(s,3,6) != "s is"{
		t.Error(Substr(s,3,6) )
	}
	s = "hello"
	if Substr(s,3,6) != "lo"{
		t.Error(Substr(s,3,6) )
	}
}

func TestSubstrLength(t *testing.T) {
	var s = "this is a chinese people"
	if SubstrLength(s,3,8) != "s is a c"{
		t.Error(SubstrLength(s,3,8))
	}
	s = "hello"
	if SubstrLength(s,3,6) != "lo"{
		t.Error(SubstrLength(s,3,8))
	}
}