package glog

import (
	"testing"
	//"fmt"
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
		Flag:LongFile,
	})
	gg.Info("this is info")
	gg.Warn("this is warn")
	gg.Error("show error")
}