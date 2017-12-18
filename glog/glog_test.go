package glog

import (
	"testing"
	//"fmt"
)

func TestAsset(t *testing.T) {
	Asset("%s","it is good")
	Asset("this %d %d",LongFile ,ShortFile)
	Info("%v",0&(0|1))
}

func TestError(t *testing.T) {
	Error("%s","error")
}

func TestAll(t *testing.T){
	//fmt.Println("test all")
	gg := NewGLogFile("../temp/1.log",&Glog{
		ShowLevel:debug,
		SaveLevel:errors,
	})
	gg.Info("INNNN")
	gg.Error("FFFF")
}