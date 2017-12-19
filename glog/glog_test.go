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
		Flag:ShortFile,
		MaxLogSize:100,
	})
	for i:=1;i<=100;i++ {
		gg.Info(i,"this is info")
		gg.Warn(i,"this is warn")
		gg.Error(i,"show error")
	}
}