package ustrings

import (
	"testing"
	//"fmt"
	"fmt"
)

func TestMatchStringNum(t *testing.T) {
	isNumber,num := MatchStringNum("123")
	if !isNumber{
		t.Error(num)
	}
	isNumber2,num2 := MatchStringNum("1a23")
	if isNumber2{
		fmt.Println(isNumber2)
		t.Error(num2)
	}
	isNumber3,num3 := MatchStringNum("aaaa")
	if isNumber3{
		fmt.Println(isNumber3)
		t.Error(num3)
	}
}