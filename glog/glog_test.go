package glog

import (
	"testing"
	//"fmt"
	"strings"
)

func TestAsset(t *testing.T) {
	Asset("%s","it is good")
	Asset("this",LongFile ,ShortFile)
	Info("bool",0&(0|1))
}

func TestError(t *testing.T) {
	Error("error","error")
}

func TestAll(t *testing.T){
	//fmt.Println("test all")
	gg := NewGLogFile("../temp/1.log",&Glog{
		ShowLevel: LevelDebug,
		SaveLevel: LevelWarn,
		Flag:ShortFile,
		MaxLogSize:1000,
	})
	for i:=1;i<=1000;i++ {
		gg.Info(i,"this is info")
		gg.Error(i,"show error",strings.Repeat(" this is a test from golang",i))
		gg.Warn(i,"this is warn")
	}
}